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
	b := models.NewBlog()

	// bs := b.GetBlogList(0, 5)

	// for item := bs.Front(); item != nil; item = item.Next() {
	// 	b, ok := item.Value.(*models.Blog)
	// 	if !ok {
	// 		panic("item not *websocket.Conn")
	// 	} else {
	// 		log.Println("Item OK")
	// 	}
	// 	bm := b.ToMap()
	// 	msg := models.NewMsg()
	// 	msg.Code = "blog"
	// 	msg.Data = b.ToJSON()
	// 	log.Println(msg.ToString())
	// 	this.Data["OneBlog"] = bm
	// 	this.Data["ArtTitle"] = bm["title"]
	// 	this.Data["ArtTime"] = bm["created"]
	// }

	bt := b.GetBlogs(0, 5)
	this.Data["Blogs"] = bt

	// this.Data["OneBlog"] = ms.ToMap()

	wp := models.NewWebPage("首页")
	wp.IncrViewCount()
	this.Data["PageTitle"] = wp.GetPageTitle()
	this.Data["ImgHost"] = wp.GetImgHost()
	this.Data["StaticHost"] = wp.GetStaticHost()
	// this.Data["TestStr"] = "s"
	this.TplNames = "index.tpl"
}
