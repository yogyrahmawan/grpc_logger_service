package api

import (
	"io"

	log "github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/yogyrahmawan/logger_service/src/pb"
	"golang.org/x/net/context"
)

// SendLog handle send log from client
func (s *Server) SendLog(srv pb.LoggerService_SendLogServer) error {
	log.Info("starting send log")
	ctx := srv.Context()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// receive data from stream
		req, err := srv.Recv()
		if err == io.EOF {
			// end of message
			log.Info("exiting because eof")
			return srv.SendAndClose(&pb.LoggerResponse{
				Status: "ok",
			})
		}

		if err != nil {
			log.Errorf("receive error %v", err)
			return err
		}

		// save it
		// TODO implements
		log.Info(req)
	}
}

// GetLog handle get log from client
func (s *Server) GetLog(ctx context.Context, in *empty.Empty) (*pb.LoggerResponsesMessage, error) {
	// TODO implements it
	return nil, nil
}
