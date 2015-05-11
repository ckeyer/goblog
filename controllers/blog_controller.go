package controllers

import (
	// "encoding/base64"
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/models"
	// "log"
	// "strconv"
	// "strings"
)

type BlogController struct {
	beego.Controller
}

func (this *BlogController) Get() {
	wp := models.NewWebPage("博客")
	wp.IncrViewCount()

	// this.Data["OneBlog"] = ms.ToMap()

	s := this.Ctx.Input.Param(":key")
	this.Ctx.WriteString(s)
	// if i, err := strconv.Atoi(s); err != nil || i == 0 {
	// 	this.checkError()
	// 	return
	// }
	// b := models.NewBlog()
	// if nil == b.ReadBlogByID(s) {
	// 	this.Data["ArticleTitle"] = b.Title
	// 	this.Data["ArticleContent"] = b.Content

	// 	// this.Data["Tags"] = b.Tags

	// 	previous := b.GetPreviousBlog()
	// 	next := b.GetNextBlog()
	// 	if previous == nil {
	// 		previous = models.NewBlog()
	// 		// previous.ID = 0
	// 		previous.Title = "6L+Z5bey5piv5pyA5YmN5LiA56+H"
	// 	}
	// 	if next == nil {
	// 		next = models.NewBlog()
	// 		// next.ID = 0
	// 		next.Title = "6L+Z5bey5piv5pyA5ZCO5LiA56+H"
	// 	}

	// 	this.Data["Previous"] = previous
	// 	this.Data["Next"] = next

	// 	// sssss := b.GetBlogsByTagId(2, 0, 5)
	// } else {
	// 	this.checkError()
	// 	return
	// }
	// this.Data["PageTitle"] = wp.GetPageTitle()
	// this.Data["ImgHost"] = wp.GetImgHost()
	// this.Data["StaticHost"] = wp.GetStaticHost()

	// this.TplNames = "blog.tpl"
}
func (this *BlogController) checkError() {
	wp := models.NewWebPage("首页")
	wp.IncrViewCount()

	this.Data["PageTitle"] = wp.GetPageTitle()
	this.Data["ImgHost"] = wp.GetImgHost()
	this.Data["StaticHost"] = wp.GetStaticHost()

	// b := models.NewBlog()
	// this.Data["Blogs"] = b.GetBlogs(0, 5)

	this.TplNames = "index.tpl"
}
func (this *BlogController) Post() {
	log.Println("")
}

// func decodeBase64(s string) string {
// 	s = strings.Replace(s, "+", "-", -1)
// 	s = strings.Replace(s, "/", "_", -1)
// 	v, _ := base64.URLEncoding.DecodeString(s)
// 	return string(v)
// }
