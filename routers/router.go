package routers

import (
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/controllers"
	"github.com/ckeyer/goblog/libs"
)

var log = libs.GetLogger()

func LoadRouters() {
	log.Info("加载路由信息")

	beego.SetStaticPath("/img", "blog/img")

	beego.Router("/", &controllers.IndexController{})
	beego.Router("/:name:string.html", &controllers.BlogController{})
	beego.Router("/tag", &controllers.TagController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/archive/:year:string-:month:string.html", &controllers.ArchiveController{})

	beego.Router("/webhook", &controllers.WebhookController{})
}
