package controllers

import (
	_ "container/list"
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/models"
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
