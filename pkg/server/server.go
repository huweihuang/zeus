package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ware "github.com/huweihuang/golib/gin/middlewares"
	log "github.com/huweihuang/golib/logger/logrus"
	"github.com/sirupsen/logrus"

	"github.com/huweihuang/gin-api-frame/cmd/server/app/config"
	"github.com/huweihuang/gin-api-frame/pkg/model"
)

type Server struct {
	conf *config.Config
	gin  *gin.Engine
}

// NewServer creates a Server object
func NewServer(conf *config.Config) *Server {
	gin.SetMode(gin.ReleaseMode)
	return &Server{
		conf: conf,
		gin:  gin.New(),
	}
}

// Run runs the specified APIServer.  This should never exit.
func (s *Server) Run() error {
	// init config
	logger := setLogger(s.conf.Log)
	if _, err := model.SetupDB(s.conf.Database); err != nil {
		return err
	}

	// setup http server
	addr := fmt.Sprintf("%s:%d", s.conf.Server.Host, s.conf.Server.Port)
	server := s.setupServer(logger)
	if s.conf.Server.CertFile != "" && s.conf.Server.KeyFile != "" {
		go server.RunTLS(addr, s.conf.Server.CertFile, s.conf.Server.KeyFile)
		log.Logger.Infof("Server listening at https://%s", addr)
	} else {
		go server.Run(addr)
		log.Logger.Infof("Server listening at http://%s", addr)
	}

	// shutting down
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := model.Close(); err != nil {
		log.Logger.Errorf("Close db error: %s", err.Error())
	}
	log.Logger.Infof("Shutting down...")
	return nil
}

func (s *Server) setupServer(logger *logrus.Logger) *gin.Engine {
	pprof.Register(s.gin)
	s.gin.Use(
		ware.RequestIDMiddleware,
		ware.LogMiddleware(logger),
		gin.RecoveryWithWriter(logger.Out),
		cors.Default(),
	)
	s.setupRoutes()
	return s.gin
}

func setLogger(conf *config.LogConfig) *logrus.Logger {
	logger := log.InitLogger(conf.LogFile, conf.LogLevel, conf.LogFormat, conf.EnableReportCaller, conf.EnableForceColors)
	return logger
}
