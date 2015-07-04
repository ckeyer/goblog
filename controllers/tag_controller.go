package controllers

import (
	// "encoding/base64"
	// "github.com/astaxie/beego"
	"github.com/ckeyer/goblog/models"
	// "log"
	// "strconv"
	// "strings"
)

type TagController struct {
	BaseController
}

func (this *TagController) Get() {

	skey := this.Ctx.Input.Param(":key")
	tag_id, err := strconv.Atoi(skey)
	if err != nil {
		this.Ctx.WriteString("Error")
		return
	}
	tag := &models.Tag{Id: tag_id}

	this.Data["TagBlogs"] = tag.GetBlogs(0, 5)
	this.TplNames = "tag.tpl"
}

// func (this *TagController) checkError() {

// 	this.TplNames = "index.tpl"
// }

// // func (this *TagController) Post() {
// // 	log.Println("")
// // }
