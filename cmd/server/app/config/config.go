package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/go-sql-driver/mysql"
)

// General configuration
type Config struct {
	Host       string
	Port       int
	Worker     int
	LogConfig  LogConfig
	Database   DBConfig
	EtcdConfig EtcdConfig
}

// LogConfig config
type LogConfig struct {
	LogFile           string
	LogLevel          string
	BackTrackLevel    string
	LogFormat         string
	EnableForceColors bool
}

// DB config
type DBConfig struct {
	User     string
	Password string
	Addr     string
	DBName   string
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

// MustLoad parse path generate the config object
func MustLoad(path string) *Config {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	conf := new(Config)
	if _, err := toml.DecodeFile(path, conf); err != nil {
		panic(err)
	}
	return conf
}
