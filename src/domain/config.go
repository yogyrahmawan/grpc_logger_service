package domain

// Config represent parameter config
type Config struct {
	RPCServer    *rpcServer
	LoggerConfig *loggerConfig
	DatabaseURL  string
}

type rpcServer struct {
	RPCHost string
	RPCPort int
}

type loggerConfig struct {
	Level string
}

// NewConfig create new configuration
func NewConfig(rpcHost, dbURL, logLevel string, rpcPort int) *Config {
	cfg := new(Config)

	// create rpcserver config
	rpcServer := new(rpcServer)
	rpcServer.RPCHost = rpcHost
	rpcServer.RPCPort = rpcPort

	cfg.RPCServer = rpcServer

	// create log
	lg := new(loggerConfig)
	lg.Level = logLevel

	cfg.LoggerConfig = lg
	cfg.DatabaseURL = dbURL

	return cfg
}
