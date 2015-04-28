package test

import (
	"github.com/ckeyer/goblog/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func init() {
	models.RegistDB()

}

func TestTag(t *testing.T) {
	Convey("should be test add tag", t, func() {
		tag := models.NewTag("TEST", -1)
		Convey("`tag` should not be nil", func() {
			So(tag, ShouldNotBeNil)
		})
		Convey("insert tag to database", func() {
			b := tag.Insert()
			So(b, ShouldEqual, true)
		})
		Convey("update tag to database", func() {
			b := tag.Update("Test", -2)
			So(b, ShouldEqual, true)
		})
		Convey("get tag by id", func() {
			_, e := models.GetTagById(tag.Id)
			So(e, ShouldEqual, nil)
		})
		Convey("delete tag to database", func() {
			b := tag.Delete()
			So(b, ShouldEqual, true)
		})
		// Convey("find tag by tagname", func() {
		// 	_, e := models.FindUserByName("sllt")
		// 	So(e, ShouldEqual, nil)
		// })
	})
}
