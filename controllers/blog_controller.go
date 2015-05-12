package controllers

import (
	// "encoding/base64"
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/models"
	// "log"
	// "strconv"
	// "strings"
)

type BlogController struct {
	beego.Controller
}

func (this *BlogController) initData() {
	this.Data["STATIC_URL_JS"] = this.GetSession("STATIC_URL_JS")
	this.Data["STATIC_URL_IMG"] = this.GetSession("STATIC_URL_IMG")
	this.Data["STATIC_URL_CSS"] = this.GetSession("STATIC_URL_CSS")
}

// @router /v1/blog [get]
func (this *BlogController) ShowList() {
	this.initData()
	log.Println("// @router /v1/blog [get]")
	wp := models.NewWebPage("博客")
	wp.IncrViewCount()
	this.Ctx.WriteString("// @router /v1/blog [get]")
}

// @router /v1/blog::key([0-9]+).html [get]
func (this *BlogController) ShowBlog() {
	this.initData()
	log.Println("// @router /v1/blog::key([0-9]+).html [get]")
	s := this.Ctx.Input.Param(":key")
	this.Ctx.WriteString(s)
	this.Ctx.WriteString("// @router /v1/blog:key [get]")
}

func (this *BlogController) Post() {
	log.Println("")
}

// @router /v1/admin/blog [get]
func (this *BlogController) ShowEditList() {
	this.initData()
	log.Println("// @router /v1/admin/blog [get]")
	this.Ctx.WriteString("// @router /v1/admin/blog [get]")
}

// @router /v1/admin/blog:key([0-9]+).html [get]
func (this *BlogController) EditBlog() {
	this.initData()
	// s := this.Ctx.Input.Param(":key")
	// this.Ctx.WriteString(s)
	// this.Ctx.WriteString("// @router /v1/admin/blog [get]")
	this.TplNames = "blogEdit.tpl"
}
