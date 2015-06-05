package test

import (
	"github.com/ckeyer/goblog/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMessage(t *testing.T) {
	Convey("should be test Message", t, func() {
		msg := &models.Message{
			Code: "code_好人啊llo",
			Data: "data_hello",
			Desc: "desc_hello",
			Supp: []string{"hello1", "hello2"},
		}
		Convey("`msg` should can MU", func() {
			bs, e := msg.ToBytes()
			So(e, ShouldBeNil)

			str := msg.ToBase64String()
			So(str, ShouldNotEqual, "")
			Println(str)

			msg2, e := models.DecodeJson([]byte(str))
			So(e, ShouldBeNil)
			So(msg2.Code, ShouldEqual, msg.Code)

			bs2, e := msg2.ToBytes()
			So(e, ShouldBeNil)
			So(string(bs2), ShouldEqual, string(bs))
		})
	})
}
