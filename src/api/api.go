package api

import (
	"fmt"
	"net"

	log "github.com/Sirupsen/logrus"

	"github.com/yogyrahmawan/logger_service/src/pb"
	"github.com/yogyrahmawan/logger_service/src/store"
	"github.com/yogyrahmawan/logger_service/src/store/mongostore"
	"github.com/yogyrahmawan/logger_service/src/utils"
	grpc "google.golang.org/grpc"
)

var (
	mongoStore store.Store
)

// RunServer initialise and run server
func RunServer() {
	st, err := initStore()
	if err != nil {
		return
	}

	mongoStore = st
	startGRPCServer()
}

func initStore() (*mongostore.MongoStore, error) {
	st, err := mongostore.NewMongoStore(utils.Cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("error at init store, err =%v", err)
		return nil, err
	}

	return st, nil
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
