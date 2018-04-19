package domain

import (
	"time"
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
