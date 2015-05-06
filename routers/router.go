package routers

import (
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/article_:id([0-9]+).html", &controllers.BlogController{})

	beego.Router("/blog", &controllers.BlogController{})
	beego.Router("/blog/:key", &controllers.BlogController{})

	beego.Router("/tag", &controllers.TagController{})
	beego.Router("/tag/:key", &controllers.TagController{})

	beego.Router("/matrix", &controllers.MatrixController{})

	beego.Router("/test", &controllers.TestController{})
}
