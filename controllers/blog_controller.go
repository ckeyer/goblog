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

// @router /v1/blog [get]
func (this *BlogController) ShowList() {
	log.Println("// @router /v1/blog [get]")
	wp := models.NewWebPage("博客")
	wp.IncrViewCount()
	this.Ctx.WriteString("// @router /v1/blog [get]")
}

// @router /v1/blog::key([0-9]+).html [get]
func (this *BlogController) ShowBlog() {
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
	log.Println("// @router /v1/admin/blog [get]")
	this.Ctx.WriteString("// @router /v1/admin/blog [get]")
}

// @router /v1/admin/blog:key([0-9]+).html [get]
func (this *BlogController) EditBlog() {
	log.Println("// @router /v1/admin/blog::key([0-9]+).html [get]")
	s := this.Ctx.Input.Param(":key")
	this.Ctx.WriteString(s)
	this.Ctx.WriteString("// @router /v1/admin/blog [get]")
}
