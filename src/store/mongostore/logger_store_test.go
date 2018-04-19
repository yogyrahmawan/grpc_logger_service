package mongostore

import (
	"testing"
	"time"

	log "github.com/Sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/yogyrahmawan/logger_service/src/domain"
	"github.com/yogyrahmawan/logger_service/src/store"
)

func TestLoggerStore(t *testing.T) {
	log.Info("execute test logger store")
	StoreTest(t, runAllTestLogger)
}

func runAllTestLogger(t *testing.T, ss store.Store) {
	t.Run("get all", func(t *testing.T) {
		getAll(t, ss)
	})

	t.Run("get by service", func(t *testing.T) {
		getByService(t, ss)
	})

	t.Run("get by level", func(t *testing.T) {
		getByLevel(t, ss)
	})
}

func getAll(t *testing.T, ss store.Store) {
	Convey("Test Get All", t, func() {
		// create logger message
		lm := domain.NewLoggerMessage("localhost:8000", "test", "debug", "sampple log", time.Now())

		Convey("empty get all", func() {
			result := <-ss.LoggerStore().GetAll()
			So(result.Err, ShouldBeNil)
			So(len(result.Data.([]*domain.LoggerMessage)), ShouldEqual, 0)
		})

		Convey("test save and then get all", func() {
			result := <-ss.LoggerStore().Save(lm)
			So(result.Err, ShouldBeNil)

			resultGet := <-ss.LoggerStore().GetAll()
			So(resultGet.Err, ShouldBeNil)

			So(len(resultGet.Data.([]*domain.LoggerMessage)), ShouldEqual, 1)
		})
	})
}

func getByService(t *testing.T, ss store.Store) {
	Convey("test by service", t, func() {
		Convey("not found", func() {
			result := <-ss.LoggerStore().GetByServiceName("not_exist")
			So(result.Err, ShouldBeNil)
			So(len(result.Data.([]*domain.LoggerMessage)), ShouldEqual, 0)
		})

		Convey("found", func() {
			result := <-ss.LoggerStore().GetByServiceName("test")
			So(result.Err, ShouldBeNil)
			So(len(result.Data.([]*domain.LoggerMessage)), ShouldEqual, 1)
		})
	})
}

func getByLevel(t *testing.T, ss store.Store) {
	Convey("test by level", t, func() {
		Convey("not found", func() {
			result := <-ss.LoggerStore().GetByLevel("not_exist")
			So(result.Err, ShouldBeNil)
			So(len(result.Data.([]*domain.LoggerMessage)), ShouldEqual, 0)
		})

		Convey("found", func() {
			result := <-ss.LoggerStore().GetByLevel("debug")
			So(result.Err, ShouldBeNil)
			So(len(result.Data.([]*domain.LoggerMessage)), ShouldEqual, 1)
		})
	})

}
