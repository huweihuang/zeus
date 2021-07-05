package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/huweihuang/gin-api-frame/cmd/server/app/config"
	"github.com/huweihuang/gin-api-frame/pkg/model"
	"github.com/huweihuang/gin-api-frame/pkg/util/log"
)

// Server is an interface for providing http service
type Server interface {
	Run() error
}

type apiserver struct {
	conf *config.Config
	gin  *gin.Engine
}

// NewAPIServer creates a Server object
func NewAPIServer(conf *config.Config) Server {
	gin.SetMode(gin.ReleaseMode)
	return &apiserver{
		conf: conf,
		gin:  gin.New(),
	}
}

// Run runs the specified APIServer.  This should never exit.
func (s *apiserver) Run() error {
	shutdown := make(chan struct{})
	setLogger(shutdown, s.conf)
	if err := setupDB(&s.conf.Database); err != nil {
		return err
	}

	addr := fmt.Sprintf("%s:%d", s.conf.Server.Host, s.conf.Server.Port)
	server := s.setupServer(log.Logger)
	if s.conf.Server.CertFile != "" && s.conf.Server.KeyFile != "" {
		go server.RunTLS(addr, s.conf.Server.CertFile, s.conf.Server.KeyFile)
		log.Logger.Infof("Server listening at https://%s", addr)
	} else {
		go server.Run(addr)
		log.Logger.Infof("Server listening at http://%s", addr)
	}

	<-shutdown

	if err := model.Close(); err != nil {
		log.Logger.Errorf("Close db error: %s", err.Error())
	}
	log.Logger.Infof("Shutting down...")
	return nil
}

func (s *apiserver) setupServer(logger *logrus.Logger) *gin.Engine {
	s.gin.Use(requestIDMiddleware, logMiddleware(log.Logger), gin.RecoveryWithWriter(log.Logger.Out), cors.Default())
	s.setupRoutes()
	return s.gin
}

func setupDB(dbConf *config.DBConfig) error {
	if err := model.SetupDB(config.FormatDSN(dbConf)); err != nil {
		return fmt.Errorf("failed to setup database")
	}
	return nil
}

func setLogger(shutdown chan struct{}, conf *config.Config) *logrus.Logger {
	logger := log.InitLogger(conf.Log.LogFile, conf.Log.LogLevel, conf.Log.BackTrackLevel, conf.Log.LogFormat,
		conf.Log.EnableForceColors)
	registerSignal(shutdown, func() {
		log.ReopenLogs(conf.Log.LogFile, logger)
	})
	return logger
}

func registerSignal(shutdown chan struct{}, logsReopenCallback func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM}...)
	go func() {
		for sig := range c {
			if handleSignals(sig, logsReopenCallback) {
				close(shutdown)
				return
			}
		}
	}()
}

func handleSignals(sig os.Signal, logsReopenCallback func()) (exitNow bool) {
	switch sig {
	case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
		return true
	case syscall.SIGUSR1:
		logsReopenCallback()
		return false
	}
	return false
}
