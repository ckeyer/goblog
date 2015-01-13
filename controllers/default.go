package controllers

import (
	"blog/models"
	_ "container/list"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	wp := models.NewWebPage("首页")
	wp.IncrViewCount()

	this.Data["PageTitle"] = wp.GetPageTitle()
	this.Data["ImgHost"] = wp.GetImgHost()
	this.Data["StaticHost"] = wp.GetStaticHost()

	b := models.NewBlog()
	this.Data["Blogs"] = b.GetBlogs(0, 5)
	this.Data["HotTags"] = b.GetHotTags()

	this.TplNames = "index.tpl"
}
