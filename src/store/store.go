package store

import (
	"github.com/yogyrahmawan/grpc_logger_service/src/domain"
)

// Result is container to get result from db
type Result struct {
	Data interface{}
	Err  *domain.ApplicationError
}

// Channel hold channel of result
type Channel chan Result

// Store is interface to interact with sql or nosql store
type Store interface {
	LoggerStore() LoggerStore
}

// LoggerStore contains logger store method
type LoggerStore interface {
	GetAll() Channel
	GetByServiceName(serviceName string) Channel
	GetByLevel(level string) Channel
	Save(*domain.LoggerMessage) Channel
}
