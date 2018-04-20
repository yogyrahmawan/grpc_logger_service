package domain

// Config represent parameter config
type Config struct {
	RPCServer    *rpcServer
	RestServer   *restServer
	LoggerConfig *loggerConfig
	DatabaseURL  string
	ServerCert   *certConfig
}

type rpcServer struct {
	RPCHost string
	RPCPort int
}

type restServer struct {
	Host string
	Port int
}

type loggerConfig struct {
	Level string
}

type certConfig struct {
	ServerCrtPath string
	ServerKeyPath string
}

// NewConfig create new configuration
func NewConfig(rpcHost, restHost, dbURL, logLevel, serverCertPath, serverKeyPath string, rpcPort, restPort int) *Config {
	cfg := new(Config)

	// create rpcserver config
	rpcServer := new(rpcServer)
	rpcServer.RPCHost = rpcHost
	rpcServer.RPCPort = rpcPort

	cfg.RPCServer = rpcServer

	// rest server
	restServer := new(restServer)
	restServer.Host = restHost
	restServer.Port = restPort
	cfg.RestServer = restServer

	// create log
	lg := new(loggerConfig)
	lg.Level = logLevel

	cfg.LoggerConfig = lg
	cfg.DatabaseURL = dbURL

	// cert
	crtConfig := new(certConfig)
	crtConfig.ServerCrtPath = serverCertPath
	crtConfig.ServerKeyPath = serverKeyPath
	cfg.ServerCert = crtConfig

	return cfg
}
