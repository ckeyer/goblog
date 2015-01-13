package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
	"log"
	"strconv"
)

type ManageController struct {
	beego.Controller
}

func (this *ManageController) Get() {
	wp := models.NewWebPage("博客")
	wp.IncrViewCount()

	// this.Data["OneBlog"] = ms.ToMap()

	s := this.Ctx.Input.Param(":key")
	if _, err := strconv.Atoi(s); err != nil {
		this.checkError()
		return
	}
	b := models.NewBlog()
	if nil == b.ReadBlogByID(s) {
		this.Data["ArticleTitle"] = b.Title
		this.Data["ArticleContent"] = b.Content

		log.Println(b.Tags)
		this.Data["Tags"] = b.Tags
		this.Data["HotTags"] = b.GetHotTags()
		for i, v := range b.GetHotTags() {
			log.Println(i, v)
		}
	} else {
		this.checkError()
		return
	}
	this.Data["PageTitle"] = wp.GetPageTitle()
	this.Data["ImgHost"] = wp.GetImgHost()
	this.Data["StaticHost"] = wp.GetStaticHost()

	this.TplNames = "manage.tpl"
}
func (this *ManageController) checkError() {
	wp := models.NewWebPage("博客-Error")

	// this.Data["OneBlog"] = ms.ToMap()

	wp.IncrViewCount()
	this.Data["PageTitle"] = wp.GetPageTitle()
	this.Data["ImgHost"] = wp.GetImgHost()
	this.Data["StaticHost"] = wp.GetStaticHost()

	this.TplNames = "blog.tpl"
}
func (this *ManageController) Post() {

}
