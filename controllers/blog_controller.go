package controllers

import (
	"blog/models"
	"github.com/astaxie/beego"
	"strconv"
)

type BlogController struct {
	beego.Controller
}

func (this *BlogController) Get() {
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
	}
	this.Data["PageTitle"] = wp.GetPageTitle()
	this.Data["ImgHost"] = wp.GetImgHost()
	this.Data["StaticHost"] = wp.GetStaticHost()

	this.TplNames = "blog.tpl"
}
func (this *BlogController) checkError() {
	wp := models.NewWebPage("博客-Error")

	// this.Data["OneBlog"] = ms.ToMap()

	wp.IncrViewCount()
	this.Data["PageTitle"] = wp.GetPageTitle()
	this.Data["ImgHost"] = wp.GetImgHost()
	this.Data["StaticHost"] = wp.GetStaticHost()

	this.TplNames = "blog.tpl"
}
func (this *BlogController) Post() {

}

// func decodeBase64(s string) string {
// 	s = strings.Replace(s, "+", "-", -1)
// 	s = strings.Replace(s, "/", "_", -1)
// 	v, _ := base64.URLEncoding.DecodeString(s)
// 	return string(v)
// }
