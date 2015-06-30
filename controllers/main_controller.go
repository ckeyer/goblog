package controllers

import (
	_ "container/list"
	"github.com/ckeyer/goblog/models"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["LatestBlogs"] = models.GetBlogs(0, 5)

	this.TplNames = "index.tpl"
}
