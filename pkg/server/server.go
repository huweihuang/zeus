package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ware "github.com/huweihuang/golib/gin/middlewares"
	log "github.com/huweihuang/golib/logger/zap"
	"go.uber.org/zap/zapcore"

	"github.com/huweihuang/zeus/cmd/server/app/configs"
	"github.com/huweihuang/zeus/pkg/client"
	"github.com/huweihuang/zeus/pkg/controller"
	"github.com/huweihuang/zeus/pkg/model"
)

type Server struct {
	conf *configs.Config
	gin  *gin.Engine
}

// NewServer creates a Server object
func NewServer(conf *configs.Config) *Server {
	gin.SetMode(gin.ReleaseMode)
	return &Server{
		conf: conf,
		gin:  gin.New(),
	}
}

// Run runs the specified APIServer.  This should never exit.
func (s *Server) Run() error {
	defer s.Shutdown()

	// init config
	err := Init(s.conf)
	if err != nil {
		return fmt.Errorf("failed to init config, err: %v", err)
	}

	// start worker controller
	workerController, err := controller.NewWorkerController(s.conf.K8s.KubeConfigPath)
	if err != nil {
		return fmt.Errorf("failed to new worker controller, err: %v", err)
	}
	go workerController.Run(context.Background(), s.conf.Worker.WorkerNumber)

	// setup http server
	addr := fmt.Sprintf("%s:%d", s.conf.Server.Host, s.conf.Server.Port)
	server := s.setupServer()
	if s.conf.Server.CertFile != "" && s.conf.Server.KeyFile != "" {
		go func() {
			err := server.RunTLS(addr, s.conf.Server.CertFile, s.conf.Server.KeyFile)
			if err != nil {
				log.Logger().With("err", err).Fatal("Failed to start http server")
			}
		}()
		log.Logger().Infof("Server listening at https://%s", addr)
	} else {
		go func() {
			err := server.Run(addr)
			if err != nil {
				log.Logger().With("err", err).Fatal("Failed to start http server")
			}
		}()
		log.Logger().Infof("Server listening at http://%s", addr)
	}

	// shutting down
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Logger().Infof("Shutting down...")
	return nil
}

func (s *Server) setupServer() *gin.Engine {
	pprof.Register(s.gin)
	s.gin.Use(
		ware.RequestIDMiddleware,
		ware.Logger(),
		gin.RecoveryWithWriter(zapcore.AddSync(os.Stdout)),
		cors.Default(),
	)
	s.setupRoutes()
	return s.gin
}

func (s *Server) Shutdown() {
	err := model.Close()
	if err != nil {
		log.Logger().Errorf("Close db error: %s", err.Error())
	}
}

func Init(conf *configs.Config) error {
	log.InitLogger(conf.Log.LogFile, conf.Log.ErrorFile, conf.Log.LogLevel, conf.Log.LogFormat, conf.Log.EnableColor)

	_, err := model.InitDB(conf.Database)
	if err != nil {
		return err
	}
	_, err = client.NewClients(conf.Client)
	if err != nil {
		return err
	}
	return nil
}
