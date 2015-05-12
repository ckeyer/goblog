package test

import (
	"fmt"
	"github.com/ckeyer/goblog/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTag(t *testing.T) {
	Convey("should be test Tag", t, func() {
		tag := models.NewTag("TEST", -1)
		Convey("`tag` should not be nil", func() {
			So(tag, ShouldNotBeNil)
		})
		Convey("insert tag to database", func() {
			e := tag.Insert()
			So(e, ShouldEqual, nil)
		})
		Convey("update tag to database", func() {
			tag.Name = "Test"
			tag.ParentId = -2
			e := tag.Update()
			So(e, ShouldEqual, nil)
		})
		Convey("get tag by id", func() {
			_, e := models.GetTagById(tag.Id)
			So(e, ShouldEqual, nil)
		})
		Convey("get tag by name", func() {
			t := models.GetTagByName("linux")
			fmt.Println(t)
			So(t.Id, ShouldNotEqual, 0)
		})
		Convey("delete tag to database", func() {
			e := tag.Delete()
			So(e, ShouldEqual, nil)
		})
	})
}
