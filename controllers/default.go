package controllers

import (
	"blog/models"
	_ "container/list"
	"github.com/astaxie/beego"
	"log"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	b := models.NewBlog()
	err := b.ReadBlogByID("10003")
	if err != nil {
		log.Println(err.Error())
	}
	bm := b.ToMap()
	this.Data["OneBlog"] = bm
	// b.Read()
	// bs := b.GetBlogList(0, 5)

	// for item := bs.Front(); item != nil; item = item.Next() {
	// 	b, ok := item.Value.(*models.Blog)
	// 	if !ok {
	// 		panic("item not *websocket.Conn")
	// 	}
	// 	bm := b.ToMap()
	// 	this.Data["OneBlog"] = bm
	this.Data["ArtTitle"] = bm["title"]
	this.Data["ArtTime"] = bm["created"]
	// }

	// this.Data["OneBlog"] = ms.ToMap()

	wp := models.NewWebPage("首页")
	wp.IncrViewCount()
	this.Data["PageTitle"] = wp.GetPageTitle()
	this.Data["img_host"] = wp.GetImgHost()

	// this.Data["TestStr"] = "s"
	this.TplNames = "index.tpl"
}
