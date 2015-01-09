package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type BlogController struct {
	beego.Controller
}

func (this *BlogController) Get() {
	wp := models.NewWebPage("首页")
	wp.IncrViewCount()

	this.Data["PageTitle"] = wp.GetPageTitle()
	this.Data["img_host"] = wp.GetImgHost()

	this.Data["PageTitle"] = this.Ctx.Input.Param(":key")

	this.TplNames = "index.tpl"
}
