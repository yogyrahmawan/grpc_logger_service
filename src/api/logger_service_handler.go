package api

import (
	"errors"

	log "github.com/Sirupsen/logrus"
	ptype "github.com/golang/protobuf/ptypes"
	"github.com/yogyrahmawan/grpc_logger_service/src/domain"
	"github.com/yogyrahmawan/grpc_logger_service/src/pb"
	"golang.org/x/net/context"
)

// Server is container for server struct
type Server struct{}

// SendLog handle send log from client
func (s *Server) SendLog(ctx context.Context, req *pb.LoggerMessage) (*pb.LoggerResponse, error) {
	// convert timestamp to time
	t, err := ptype.Timestamp(req.GetCreatedAt())
	if err != nil {
		log.Errorf("error converting timestamp, err = %v", err)
		return nil, err
	}

	logMsg := domain.NewLoggerMessage(req.GetIpPort(),
		req.GetServiceName(),
		req.GetLevel(),
		req.GetText(),
		t)
	saveRes := <-mongoStore.LoggerStore().Save(logMsg)
	if saveRes.Err != nil {
		log.Errorf("error when saving data, err =%v", saveRes.Err.Error())
		return nil, errors.New(saveRes.Err.Error())
	}

	return &pb.LoggerResponse{
		Status: "ok",
	}, nil

}

// GetLog handle get log from client
func (s *Server) GetLog(ctx context.Context, in *pb.GetLoggerRequest) (*pb.LoggerResponsesMessage, error) {
	// request type : all, service, level
	requestType := getRequestType(in)
	switch requestType {
	case "level":
		res, err := handleGetLevel(in.GetLevel())
		if err != nil {
			log.Errorf("Error at handle level, err = %v", err)
			return nil, err
		}
		return res, nil
	case "service_name":
		res, err := handleGetServiceName(in.GetServiceName())
		if err != nil {
			log.Errorf("error at handle service name, err = %v", err)
			return nil, err
		}
		return res, nil
	default:
		res, err := handleGetAll()
		if err != nil {
			log.Errorf("error handle get all, err = %v", err)
			return nil, err
		}
		return res, nil
	}

}

func handleGetLevel(level string) (*pb.LoggerResponsesMessage, error) {
	res := <-mongoStore.LoggerStore().GetByLevel(level)
	if res.Err != nil {
		return nil, errors.New(res.Err.Error())
	}

	return domain.LoggerMessagesToLoggerResponses(res.Data.([]*domain.LoggerMessage))
}

func handleGetServiceName(serviceName string) (*pb.LoggerResponsesMessage, error) {
	res := <-mongoStore.LoggerStore().GetByServiceName(serviceName)
	if res.Err != nil {
		return nil, errors.New(res.Err.Error())
	}

	return domain.LoggerMessagesToLoggerResponses(res.Data.([]*domain.LoggerMessage))
}

func handleGetAll() (*pb.LoggerResponsesMessage, error) {
	res := <-mongoStore.LoggerStore().GetAll()
	if res.Err != nil {
		return nil, errors.New(res.Err.Error())
	}

	return domain.LoggerMessagesToLoggerResponses(res.Data.([]*domain.LoggerMessage))
}

func getRequestType(in *pb.GetLoggerRequest) string {
	if in.GetLevel() != "" {
		return "level"
	}

	if in.GetServiceName() != "" {
		return "service_name"
	}

	return "all"
}
