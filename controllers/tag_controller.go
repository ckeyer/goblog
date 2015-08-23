package controllers

import (
	// "encoding/base64"
	// "github.com/astaxie/beego"
	"github.com/ckeyer/goblog/models"
	// "log"
	"strconv"
	// "strings"
)

type TagController struct {
	BaseController
}

func (t *TagController) Get() {

	skey := t.Ctx.Input.Param(":key")
	tag_id, err := strconv.Atoi(skey)
	if err != nil {
		t.Ctx.WriteString("Error")
		return
	}
	tag := &models.Tag{Id: int64(tag_id)}

	t.Data["TagBlogs"] = tag.GetBlogs(0, 5)
	t.TplNames = "tag.tpl"
}

// func (this *TagController) checkError() {

// 	t.TplNames = "index.tpl"
// }

func (t *TagController) Post() {
	t.Ctx.WriteString("hello post")
}
