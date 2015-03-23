package routers

import (
	"blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/blog", &controllers.BlogController{})
	beego.Router("/blog/:key", &controllers.BlogController{})

	beego.Router("/tag", &controllers.TagController{})
	beego.Router("/tag/:key", &controllers.TagController{})
	beego.Router("/test", &controllers.TestController{})
}
