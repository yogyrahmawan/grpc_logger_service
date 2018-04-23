package mongostore

import (
	"github.com/yogyrahmawan/grpc_logger_service/src/domain"
	"github.com/yogyrahmawan/grpc_logger_service/src/store"
	"gopkg.in/mgo.v2/bson"
)

// MongoLoggerStore hold struct from mongostore
type MongoLoggerStore struct {
	*MongoStore
}

// NewNoSQLLoggerStore instantiate logger store
func NewNoSQLLoggerStore(noSQLStore *MongoStore) store.LoggerStore {
	return &MongoLoggerStore{noSQLStore}
}

// GetAll get all
func (m *MongoLoggerStore) GetAll() store.Channel {
	channel := make(store.Channel)
	go func() {
		result := store.Result{}
		sesCopy := m.session.Copy()
		defer sesCopy.Close()

		var loggers []*domain.LoggerMessage
		if err := m.dB(sesCopy).C("logger").Find(nil).All(&loggers); err != nil {
			result.Err = domain.NewStoreError("Logger.GetAll",
				"Error get all records",
				"error = "+err.Error(),
			)
			channel <- result
			close(channel)
			return
		}

		result.Data = loggers
		channel <- result
		close(channel)
	}()
	return channel
}

// Save save logger
func (m *MongoLoggerStore) Save(lm *domain.LoggerMessage) store.Channel {
	channel := make(store.Channel)

	go func() {
		sesCopy := m.session.Copy()
		defer sesCopy.Close()
		result := store.Result{}

		if err := m.dB(sesCopy).C("logger").Insert(lm); err != nil {
			result.Err = domain.NewStoreError("MongoLoggerStore.SaveData",
				"error saving data",
				"detail : "+err.Error(),
			)

			channel <- result
			close(channel)
		}

		channel <- result
		close(channel)
	}()

	return channel
}

// GetByServiceName get by service name
func (m *MongoLoggerStore) GetByServiceName(serviceName string) store.Channel {
	channel := make(store.Channel)

	go func() {
		sesCopy := m.session.Copy()
		defer sesCopy.Close()
		result := store.Result{}

		loggers := []*domain.LoggerMessage{}
		if err := m.dB(sesCopy).C("logger").Find(bson.M{"service_name": serviceName}).All(&loggers); err != nil {
			result.Err = domain.NewStoreError(
				"MongoLoggerStore.GetByServiceName",
				"error getByServiceName",
				"error "+err.Error(),
			)
			channel <- result
			close(channel)
			return
		}

		result.Data = loggers
		channel <- result
		close(channel)
	}()

	return channel
}

// GetByLevel get by level
func (m *MongoLoggerStore) GetByLevel(level string) store.Channel {
	channel := make(store.Channel)

	go func() {
		sesCopy := m.session.Copy()
		defer sesCopy.Close()

		result := store.Result{}
		var loggers []*domain.LoggerMessage

		if err := m.dB(sesCopy).C("logger").Find(bson.M{"level": level}).All(&loggers); err != nil {
			result.Err = domain.NewStoreError(
				"MongoLoggerStore.GetByServiceName",
				"error getByServiceName",
				"error "+err.Error(),
			)
			channel <- result
			close(channel)
			return
		}

		result.Data = loggers
		channel <- result
		close(channel)

	}()

	return channel
}
