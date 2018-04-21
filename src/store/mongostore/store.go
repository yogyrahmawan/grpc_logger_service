package mongostore

import (
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/yogyrahmawan/grpc_logger_service/src/store"
	"gopkg.in/mgo.v2"
)

// MongoStore hold store and db connection
type MongoStore struct {
	dbName      string
	session     *mgo.Session
	loggerStore store.LoggerStore
}

// NewMongoStore create new mongo store
func NewMongoStore(dbURL string) (*MongoStore, error) {
	mongoStore := new(MongoStore)

	// setup connection
	ses, err := setupDB(dbURL)
	if err != nil {
		return nil, err
	}

	mongoStore.session = ses
	mongoStore.dbName = getDbName(dbURL)
	mongoStore.loggerStore = NewNoSQLLoggerStore(mongoStore)
	return mongoStore, nil
}

func setupDB(dbURL string) (*mgo.Session, error) {
	session, err := mgo.Dial(dbURL)
	if err != nil {
		log.Fatal("cant connect to mongo, err = " + err.Error())
		return nil, err
	}

	// set mode
	session.SetMode(mgo.Monotonic, true)

	session.SetSafe(&mgo.Safe{})

	return session, nil
}

// e.g mongodb://localhost/test?replicaSet=test
func getDbName(dbURL string) string {
	uriParts := strings.SplitN(dbURL, "/", 4)

	// get the last part (maybe a name with or without options)
	name := uriParts[len(uriParts)-1]

	// get the first part of name (omit options if any)
	nameParts := strings.SplitN(name, "?", 2)
	return nameParts[0]
}

func (m *MongoStore) dB(localSession *mgo.Session) *mgo.Database {
	return localSession.DB(m.dbName)
}

// LoggerStore hold logger store
func (m *MongoStore) LoggerStore() store.LoggerStore {
	return m.loggerStore
}
