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
	log.Debug("hello Debug")
	log.Info("hello Info")
	log.Notice("hello Notice")
	log.Warning("hello Warning")
	log.Error("hello Error")
	this.TplNames = "index.tpl"
}
