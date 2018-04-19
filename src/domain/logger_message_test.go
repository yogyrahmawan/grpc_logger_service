package domain

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLoggerMessage(t *testing.T) {
	Convey("test logger message", t, func() {
		lm := NewLoggerMessage("10.18.80.1:8080", "ini_service", "debug", "logging", time.Now())
		Convey("testify the struct", func() {
			So(lm.IPPort, ShouldEqual, "10.18.80.1:8080")
			So(lm.ServiceName, ShouldEqual, "ini_service")
			So(lm.Level, ShouldEqual, "debug")
			So(lm.Text, ShouldEqual, "logging")
		})
	})
}
