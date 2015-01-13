package controllers

import (
	"blog/models"
	_ "container/list"
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (this *TestController) Get() {

	wp := models.NewWebPage("测试")
	wp.IncrViewCount()
	this.Data["PageTitle"] = wp.GetPageTitle()
	this.Data["ImgHost"] = wp.GetImgHost()
	this.Data["StaticHost"] = wp.GetStaticHost()
	// this.Data["TestStr"] = "s"
	this.TplNames = "test.tpl"
}
