package controllers

import (
	_ "container/list"
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/models"
	logpkg "log"
	"os"
)

var log *logpkg.Logger

func init() {
	log = logpkg.New(os.Stderr, "controller", logpkg.Ltime|logpkg.Lshortfile)
}

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	wp := models.NewWebPage("首页")
	wp.IncrViewCount()

	this.Data["PageTitle"] = wp.GetPageTitle()
	this.Data["ImgHost"] = wp.GetImgHost()
	this.Data["StaticHost"] = wp.GetStaticHost()

	this.Data["Tags"] = models.GetHotTags(5)
	// this.Data["HotTags"] = b.GetHotTags()

	this.TplNames = "index.tpl"
}
