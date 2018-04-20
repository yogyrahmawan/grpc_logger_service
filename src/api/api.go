package api

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"

	"github.com/yogyrahmawan/logger_service/src/pb"
	"github.com/yogyrahmawan/logger_service/src/store"
	"github.com/yogyrahmawan/logger_service/src/store/mongostore"
	"github.com/yogyrahmawan/logger_service/src/utils"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
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

	go func() {
		if err := startGRPCServer(); err != nil {
			log.Fatalf("failed start grpc, err = %v", err)
		}
	}()

	go func() {
		if err := startRestAPIServer(); err != nil {
			log.Fatalf("failed start rest, err = %v", err)
		}
	}()

	// block indifinitely
	select {}
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

	// create tls
	creds, err := credentials.NewServerTLSFromFile(utils.Cfg.ServerCert.ServerCrtPath,
		utils.Cfg.ServerCert.ServerKeyPath)
	if err != nil {
		log.Errorf("could not load tls, err = %v", err)
		return err
	}

	// create grpc options
	opts := []grpc.ServerOption{grpc.Creds(creds),
		grpc.UnaryInterceptor(authUnaryInterceptor),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterLoggerServiceServer(grpcServer, &s)

	log.Info("Starting grpc server on %v", utils.Cfg.RPCServer.RPCPort)
	if err := grpcServer.Serve(listen); err != nil {
		log.Errorf("failed to start GRPC server, err = %v", err)
		return err
	}

	return nil
}

func startRestAPIServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// rest place token in the http header
	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(credMatcher))

	// get certificate
	creds, err := credentials.NewClientTLSFromFile(utils.Cfg.ServerCert.ServerCrtPath, "")
	if err != nil {
		log.Errorf("cannot load tls, err = %v", err)
		return err
	}

	// setup options
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	// RegisterService
	err = pb.RegisterLoggerServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", utils.Cfg.RPCServer.RPCPort), opts)
	if err != nil {
		log.Errorf("cannot register service : %v", err)
		return err
	}

	log.Info("starting HTTP 1 prtocol on " + fmt.Sprintf(":%d", utils.Cfg.RestServer.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", utils.Cfg.RPCServer.RPCPort), mux)

	return nil
}

func credMatcher(header string) (metadataName string, ok bool) {
	if header == "token" {
		return header, true
	}

	return "", false
}

// client token use JWT with hmac algorithm
func validateJWTToken(ctx context.Context) error {
	if mtd, ok := metadata.FromIncomingContext(ctx); ok {
		clientToken := strings.Join(mtd["token"], "")
		if !claimJWTToken(clientToken) {
			return fmt.Errorf("not valid token")
		}

		return nil
	}

	return errors.New("missing credentials")
}

func claimJWTToken(clientToken string) bool {
	log.Info("got token : " + clientToken)
	token, err := jwt.Parse(clientToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("113070"), nil
	})

	if err != nil {
		log.Errorf("error when parsing jwt, err = %v", err)
		return false
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// TODO handle this
		return true
	}

	return false
}

func authUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	_, ok := info.Server.(*Server)
	if !ok {
		return nil, errors.New("unable to cast server")
	}
	err := validateJWTToken(ctx)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}
