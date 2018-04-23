package domain

import . "github.com/smartystreets/goconvey/convey"
import "testing"

func TestStoreError(t *testing.T) {
	Convey("test application error ", t, func() {
		// creating object
		n := NewStoreError("this", "error", "error construct")

		Convey("validate creating object", func() {
			So(n.At, ShouldEqual, "this")
			So(n.Message, ShouldEqual, "error")
			So(n.Details, ShouldEqual, "error construct")
		})

		Convey("construct error", func() {
			So(n.Error(), ShouldContainSubstring, "error construct")
		})
	})
}
