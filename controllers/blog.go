package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type BlogController struct {
	beego.Controller
}

func (this *BlogController) Get() {
	wp := models.NewWebPage("博客")
	wp.IncrViewCount()

	this.Data["PageTitle"] = wp.GetPageTitle() + this.Ctx.Input.Param(":key")
	this.Data["img_host"] = wp.GetImgHost()

	this.TplNames = "blog.tpl"
}
