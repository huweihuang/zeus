package config

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"

	log "github.com/huweihuang/golib/logger/logrus"
)

var ApiConfig *Config

// Config is general configuration
type Config struct {
	Server   *ServerConfig
	Log      *LogConfig
	Database *DBConfig
	Etcd     *EtcdConfig
}

// ServerConfig is http server config
type ServerConfig struct {
	Host     string
	Port     int
	CertFile string
	KeyFile  string
}

// LogConfig is config for logger
type LogConfig struct {
	LogFile            string
	LogLevel           string
	LogFormat          string
	EnableReportCaller bool
	EnableForceColors  bool
}

// DBConfig is config for db
type DBConfig struct {
	User     string
	Password string
	Addr     string
	DBName   string
	LogLevel string
}

// Etcd config
type EtcdConfig struct {
	Endpoints   string
	CertFile    string
	KeyFile     string
	CAFile      string
	JobQueueKey string
}

// FormatDSN formats the given Config into a DSN string which can be passed to the driver.
func FormatDSN(dbConf *DBConfig) string {
	cfg := mysql.Config{
		User:                 dbConf.User,
		Passwd:               dbConf.Password,
		Net:                  "tcp",
		Addr:                 dbConf.Addr,
		DBName:               dbConf.DBName,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	return cfg.FormatDSN()
}

func InitConfig(configPath string) (*Config, error) {
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read in config by viper, err: %v", err)
	}

	err := viper.Unmarshal(ApiConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal, err: %v", err)
	}
	log.Logger.WithField("config", ApiConfig).Debug("init config")

	return ApiConfig, nil
}
