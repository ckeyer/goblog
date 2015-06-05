package test

import (
	"github.com/ckeyer/goblog/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestBlog(t *testing.T) {
	Convey("should be test Blog", t, func() {
		blog := &models.Blog{
			Title:   "Test Blog",
			Summary: "Test Blog content ",
			Content: "something test",
		}
		Convey("`blog` should not be nil", func() {
			So(blog, ShouldNotBeNil)
		})
		Convey("insert blog to database", func() {
			e := blog.Insert()
			So(e, ShouldEqual, nil)
		})
		Convey("update blog to database", func() {
			blog.Insert()
			blog.Title = "Test Blog(new)"
			blog.Status = -1
			e := blog.Update()
			So(e, ShouldEqual, nil)
		})
		Convey("get blog by id", func() {
			blog.Insert()
			newblog, e := models.GetBlogById(blog.Id)
			So(e, ShouldEqual, nil)
			So(newblog.Title, ShouldEqual, blog.Title)
		})
		// Convey("delete blog to database", func() {
		// 	blog.Insert()
		// 	e := blog.Delete()
		// 	So(e, ShouldEqual, nil)
		// })
	})
}
