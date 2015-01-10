package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	b := models.NewBlog()
	_ = b
	b.Read()
	defer b.Close()
	wp := models.NewWebPage("首页")
	wp.IncrViewCount()

	this.Data["PageTitle"] = wp.GetPageTitle()
	this.Data["img_host"] = wp.GetImgHost()

	// this.Data["TestStr"] = "s"
	this.TplNames = "index.tpl"
}
