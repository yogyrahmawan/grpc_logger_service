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

func TestConvertLoggerMessagesToProtoLoggerResponses(t *testing.T) {
	Convey("test convert logger messages to protologger responses", t, func() {
		var loggerMessages []*LoggerMessage
		firstData := NewLoggerMessage("localhost:8000", "test", "debug", "sample log", time.Now())
		secondData := NewLoggerMessage("localhost:8000", "test", "debug", "sample log 2", time.Now())

		loggerMessages = append(loggerMessages, firstData)
		loggerMessages = append(loggerMessages, secondData)

		Convey("validate field", func() {
			data, err := LoggerMessagesToLoggerResponses(loggerMessages)
			So(err, ShouldBeNil)
			So(len(data.LoggerMessages), ShouldEqual, 2)

		})
	})
}
