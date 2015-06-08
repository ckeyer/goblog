package test

import (
	"fmt"
	"github.com/ckeyer/goblog/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestTag(t *testing.T) {
	fmt.Println("Start")
	models.RegistDB()
	Convey("should be test Tag", t, func() {

		Convey("get hot tags", func() {
			tags := models.GetHotTags(3)
			So(len(tags), ShouldEqual, 3)
		})
		Convey("get blog count", func() {
			tags := models.GetHotTags(3)
			count := tags[0].GetBlogCount()
			So(count, ShouldBeGreaterThan, 0)
		})
		Convey("get blogs", func() {

			tags := models.GetHotTags(3)
			tag := &models.Tag{Id: tags[0].Id}
			So(len(tag.Blogs), ShouldEqual, 0)
			tag.Get()
			So(len(tag.Blogs), ShouldBeGreaterThan, 0)
		})
		// tag := models.NewTag("TEST", -1)
		// Convey("`tag` should not be nil", func() {
		// 	So(tag, ShouldNotBeNil)
		// })
		// Convey("update tag to database", func() {
		// 	tag.Name = "Test"
		// 	tag.ParentId = -2
		// 	e := tag.Update()
		// 	So(e, ShouldEqual, nil)
		// })
		// Convey("get tag by name", func() {
		// 	t := models.GetTag("linux")
		// 	fmt.Println("new tag", t)
		// 	So(t.Id, ShouldNotEqual, 0)
		// })
		// Convey("delete tag to database", func() {
		// 	e := tag.Delete()
		// 	So(e, ShouldEqual, nil)
		// })
	})
}
