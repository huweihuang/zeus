package log

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// Public logger
var (
	Logger     *logrus.Logger
	HttpLogger *logrus.Entry
)

const (
	defaultLevel          = "info"
	defaultBacktrackLevel = "warn"
)

// Reopen log fd handlers when receiving signal syscall.SIGUSR1
func ReopenLogs(logFile string, logger *logrus.Logger) error {
	if logFile == "" {
		return nil
	}
	accessLog, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	oldFd := logger.Out.(*os.File)
	logger.Out = accessLog
	oldFd.Close()

	return nil
}

// @backtrackLevel: log the backtrack info when logging level is >= backtrackLevel
func InitLogger(logFile, logLevel, backtrackLevel, format string, enableForceColors bool) *logrus.Logger {
	logger := logrus.New()

	// set log level
	if logLevel == "" {
		logLevel = defaultLevel
	}
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic("Failed to parse log level")
	}
	logger.SetLevel(level)

	// set stdout
	logger.SetOutput(os.Stdout)
	// set logfile if not empty
	if logFile != "" {
		lastIdx := strings.LastIndexAny(logFile, "/")
		err := os.MkdirAll(logFile[:lastIdx], 644)
		if err != nil {
			panic("Failed to create log directory")
		}
		// accessLog, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		accessLog, err := rotatelogs.New(
			logFile+".%Y%m%d",
			rotatelogs.WithLinkName(logFile),
			rotatelogs.WithMaxAge(time.Duration(7*24)*time.Hour),
			rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		)
		if err != nil {
			panic("Failed to create access.log")
		}
		writers := []io.Writer{
			accessLog,
			os.Stdout,
		}
		fileAndStdoutWriter := io.MultiWriter(writers...)
		logger.SetOutput(fileAndStdoutWriter)
	}

	forceColors := false
	if enableForceColors {
		forceColors = true
	}
	// set file && line number
	logger.SetReportCaller(true)
	switch format {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcName := s[len(s)-1]
				fileName := path.Base(f.File)
				return funcName, fmt.Sprintf("%s:%d", fileName, f.Line)
			},
		})
	default:
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			ForceColors:     forceColors,
			DisableQuote:    true,
			TimestampFormat: "2006-01-02 15:04:05",
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				s := strings.Split(f.Function, ".")
				funcName := s[len(s)-1]
				fileName := path.Base(f.File)
				return funcName, fmt.Sprintf("%s:%d", fileName, f.Line)
			},
		})
	}

	// set BackTrackHook
	if backtrackLevel == "" {
		backtrackLevel = defaultBacktrackLevel
	}
	btLevel, err := logrus.ParseLevel(backtrackLevel)
	if err != nil {
		panic("Failed to parser backtrack level")
	}
	logger.Hooks.Add(NewBackTrackHook(btLevel))

	Logger = logger
	return logger
}

// 初始化http logger
func InitHttpLogger(c *gin.Context) *logrus.Entry {
	reqID := c.GetString("req_id")
	if reqID == "" {
		return logrus.NewEntry(Logger)
	}
	logger := Logger.WithField("req_id", reqID)
	HttpLogger = logger
	return logger
}
