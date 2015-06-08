package controllers

import (
	_ "container/list"
	"github.com/astaxie/beego"
	// "github.com/ckeyer/goblog/models"
)

type TestController struct {
	beego.Controller
}

func (this *TestController) initData() {
	this.Data["STATIC_URL_JS"] = this.GetSession("STATIC_URL_JS")
	this.Data["STATIC_URL_IMG"] = this.GetSession("STATIC_URL_IMG")
	this.Data["STATIC_URL_CSS"] = this.GetSession("STATIC_URL_CSS")
}

func (this *TestController) Get() {
	this.initData()
	this.Data["PageTitle"] = "测试"
	// this.Data["TestStr"] = "s"
	this.TplNames = "404.tpl"
}
