package controllers

import (
	_ "container/list"
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["LatestBlogs"] = models.GetBlogs(0, 5)

	this.TplNames = "index.tpl"
}
