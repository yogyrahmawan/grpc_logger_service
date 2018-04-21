package main

import (
	"log"

	"github.com/yogyrahmawan/grpc_logger_service/cmd"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
