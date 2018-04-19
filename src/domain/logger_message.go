package domain

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/yogyrahmawan/logger_service/src/pb"
)

// LoggerMessage is internal logger message struct
type LoggerMessage struct {
	IPPort      string    `bson:"ip_port"`
	ServiceName string    `bson:"service_name"`
	Level       string    `bson:"level"`
	Text        string    `bson:"text"`
	CreatedAt   time.Time `bson:"created_at"`
}

// NewLoggerMessage create logger message
func NewLoggerMessage(ipPort, serviceName, level, text string, createdAt time.Time) *LoggerMessage {
	lm := new(LoggerMessage)
	lm.IPPort = ipPort
	lm.ServiceName = serviceName
	lm.Level = level
	lm.Text = text
	lm.CreatedAt = createdAt

	return lm
}

// LoggerMessagesToLoggerResponses convert slice of loggerMessage to loggerResponse
func LoggerMessagesToLoggerResponses(lms []*LoggerMessage) (*pb.LoggerResponsesMessage, error) {
	var pbs []*pb.LoggerMessage

	semChan := make(chan *pb.LoggerMessage, len(lms))
	errChan := make(chan error)
	defer func() {
		close(semChan)
		close(errChan)
	}()
	for _, l := range lms {
		go func(v *LoggerMessage) {
			pbLRM := new(pb.LoggerMessage)
			pbLRM.IpPort = v.IPPort
			pbLRM.ServiceName = v.ServiceName
			pbLRM.Level = v.Level
			pbLRM.Text = v.Text

			// convert to proto
			cvtTimestamp, err := ptypes.TimestampProto(v.CreatedAt)
			if err != nil {
				errChan <- err
			}

			pbLRM.CreatedAt = cvtTimestamp

			semChan <- pbLRM
		}(l)
	}

	for i := 0; i < len(lms); i++ {
		select {
		case pblrm := <-semChan:
			pbs = append(pbs, pblrm)
		case err := <-errChan:
			return nil, err
		}
	}

	return &pb.LoggerResponsesMessage{
		LoggerMessages: pbs,
	}, nil
}
