package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
)

type BlogController struct {
	beego.Controller
}

func (this *BlogController) Get() {
	wp := models.NewWebPage("博客")

	// this.Data["OneBlog"] = ms.ToMap()

	wp.IncrViewCount()
	this.Data["PageTitle"] = wp.GetPageTitle()
	this.Data["ImgHost"] = wp.GetImgHost()
	this.Data["StaticHost"] = wp.GetStaticHost()

	this.TplNames = "blog.tpl"
}
func (this *BlogController) checkError() {
	wp := models.NewWebPage("博客")

	// this.Data["OneBlog"] = ms.ToMap()

	wp.IncrViewCount()
	this.Data["PageTitle"] = wp.GetPageTitle()
	this.Data["ImgHost"] = wp.GetImgHost()
	this.Data["StaticHost"] = wp.GetStaticHost()

	this.TplNames = "blog.tpl"
}
func (this *BlogController) Post() {

}
