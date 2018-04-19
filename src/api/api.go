package api

import (
	"fmt"
	"net"

	log "github.com/Sirupsen/logrus"
	"github.com/yogyrahmawan/logger_service/src/utils"
	"github.com/yogyrahmawan/logger_service/src/pb"
	grpc "google.golang.org/grpc"
)

// Server is container for server struct
type Server struct{}

// RunServer run rpc server
func RunServer() {
	startGRPCServer()
}

func startGRPCServer() error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", utils.Cfg.RPCServer.RPCPort))
	if err != nil {
		log.Errorf("net failed to listen, err : %v", err)
		return err
	}

	// define new server
	s := Server{}
	grpcServer := grpc.NewServer()
	pb.RegisterLoggerServiceServer(grpcServer, &s)

	log.Info("Starting grpc server on %v", utils.Cfg.RPCServer.RPCPort)
	if err := grpcServer.Serve(listen); err != nil {
		log.Errorf("failed to start GRPC server, err = %v", err)
		return err
	}

	return nil
}
