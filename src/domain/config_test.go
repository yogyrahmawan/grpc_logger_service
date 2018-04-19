package domain

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHelloWorld(t *testing.T) {
	Convey("Test create config", t, func() {
		Convey("test config", func() {
			cfg := NewConfig("localhost", "postgres:postgres", "debug", 8080)
			So(cfg.RPCServer.RPCHost, ShouldEqual, "localhost")
			So(cfg.RPCServer.RPCPort, ShouldEqual, 8080)
			So(cfg.LoggerConfig.Level, ShouldEqual, "debug")
			So(cfg.DatabaseURL, ShouldEqual, "postgres:postgres")
		})
	})
}
