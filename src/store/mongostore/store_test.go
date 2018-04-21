package mongostore

import (
	"os"
	"testing"

	log "github.com/Sirupsen/logrus"
	"github.com/yogyrahmawan/grpc_logger_service/src/store"
	"github.com/yogyrahmawan/grpc_logger_service/src/store/storetest"
)

var testingStore = struct {
	Container *storetest.RunningContainer
	Store     store.Store
}{}

// StoreTest is wrapper to run bdd test
func StoreTest(t *testing.T, f func(*testing.T, store.Store)) {
	t.Run("mongo", func(t *testing.T) {
		f(t, testingStore.Store)
	})
}

func initStores() {
	container, dbURL, err := storetest.NewMongoDBContainer()
	if err != nil {
		log.Error("error init stores, err = " + err.Error())
		return
	}

	testingStore.Container = container
	testingStore.Store, err = NewMongoStore(dbURL)
	if err != nil {
		log.Error("error creating mongostore in init stores, err = " + err.Error())
		return
	}
}

func stopStores() {
	testingStore.Container.Stop()
}

func TestMain(m *testing.M) {
	initStores()
	status := 0
	defer func() {
		stopStores()
		os.Exit(status)
	}()

	status = m.Run()
}
