package cmd

import (
	"github.com/yogyrahmawan/logger_service/src/api"
	"github.com/yogyrahmawan/logger_service/src/utils"

	"github.com/spf13/cobra"
)

var rootCommand = cobra.Command{
	Use: "iso_converter",
	Run: run,
}

// RootCommand create root command
func RootCommand() *cobra.Command {
	rootCommand.PersistentFlags().StringP("config", "c", "", "the config file to use")
	rootCommand.PersistentFlags().StringP("env", "e", "", "environment : dev, test,prod")

	return &rootCommand
}

func run(cmd *cobra.Command, args []string) {
	utils.LoadConfig(cmd)
	api.RunServer()
}
