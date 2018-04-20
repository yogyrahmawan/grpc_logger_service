package utils

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yogyrahmawan/logger_service/src/domain"
)

var (
	// Cfg store configuration
	Cfg *domain.Config
)

// LoadConfig load configuration from config file
func LoadConfig(cmd *cobra.Command) {
	var environment string

	err := viper.BindPFlags(cmd.Flags())
	if err != nil {
		log.Fatalf("error bind flags, err = %v", err)
		return
	}

	if configFile, _ := cmd.Flags().GetString("config"); configFile != "" {
		log.Println("read config in file = " + configFile)
		viper.SetConfigFile(configFile)
	} else {
		log.Println("no config file. use default path")
		viper.SetConfigName("config")
		viper.AddConfigPath("./")
		viper.AddConfigPath("$HOME/.toml")
	}

	if env, _ := cmd.Flags().GetString("env"); env != "" {
		log.Println("running server at environment, env = " + env)
		environment = env
	} else {
		environment = "test"
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read config, err = " + err.Error())
		return
	}

	cfg := domain.NewConfig(viper.GetString(environment+".rpc_host"),
		viper.GetString(environment+".rest_host"),
		viper.GetString(environment+".database_url"),
		viper.GetString(environment+".log_level"),
		viper.GetString(environment+".server_cert"),
		viper.GetString(environment+".server_key"),
		viper.GetInt(environment+".rpc_port"),
		viper.GetInt(environment+".rest_port"),
	)

	log.Printf("config : %v", cfg)

	Cfg = cfg
}
