package test

// import (
// 	"fmt"
// 	"github.com/ckeyer/goblog/models"
// 	. "github.com/smartystreets/goconvey/convey"
// 	"testing"
// )

// func TestTag(t *testing.T) {
// 	fmt.Println("Start")
// 	Convey("should be test Tag", t, func() {
// 		tag := models.NewTag("TEST", -1)
// 		Convey("`tag` should not be nil", func() {
// 			So(tag, ShouldNotBeNil)
// 		})
// 		Convey("update tag to database", func() {
// 			tag.Name = "Test"
// 			tag.ParentId = -2
// 			e := tag.Update()
// 			So(e, ShouldEqual, nil)
// 		})
// 		Convey("get tag by name", func() {
// 			t := models.GetTag("linux")
// 			fmt.Println("new tag", t)
// 			So(t.Id, ShouldNotEqual, 0)
// 		})
// 		Convey("delete tag to database", func() {
// 			e := tag.Delete()
// 			So(e, ShouldEqual, nil)
// 		})
// 	})
// }
