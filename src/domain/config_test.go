package domain

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHelloWorld(t *testing.T) {
	Convey("Test create config", t, func() {
		Convey("test config", func() {
			cfg := NewConfig("localhost", "localhost", "postgres:postgres", "debug", "./cert/.crt", "./cert/.key", 8080, 8081)
			So(cfg.RPCServer.RPCHost, ShouldEqual, "localhost")
			So(cfg.RPCServer.RPCPort, ShouldEqual, 8080)
			So(cfg.RestServer.Host, ShouldEqual, "localhost")
			So(cfg.RestServer.Port, ShouldEqual, 8081)
			So(cfg.LoggerConfig.Level, ShouldEqual, "debug")
			So(cfg.DatabaseURL, ShouldEqual, "postgres:postgres")
			So(cfg.ServerCert.ServerCrtPath, ShouldEqual, "./cert/.crt")
			So(cfg.ServerCert.ServerKeyPath, ShouldEqual, "./cert/.key")
		})
	})
}
