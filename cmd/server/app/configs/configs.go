package configs

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

// EtcdConfig is config for etcd
type EtcdConfig struct {
	Endpoints   string
	CertFile    string
	KeyFile     string
	CAFile      string
	JobQueueKey string
}
