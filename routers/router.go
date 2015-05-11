package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/ckeyer/goblog/controllers"
)

func init() {
	// beego.Router("/", &controllers.MainController{})

	// beego.Router("/article_:id([0-9]+).html", &controllers.BlogController{})

	// beego.Router("/blog", &controllers.BlogController{})
	// beego.Router("/blog/:key([0-9]+).html", &controllers.BlogController{})
	// beego.Router("/tag", &controllers.TagController{})
	// beego.Router("/tag/:key", &controllers.TagController{})

	// beego.Router("/matrix", &controllers.MatrixController{})

	// beego.Router("/test", &controllers.TestController{})

	ns :=
		beego.NewNamespace("-",
			//此处正式版时改为验证加密请求
			beego.NSCond(func(ctx *context.Context) bool {
				println("hello all")
				if ua := ctx.Input.Request.UserAgent(); ua != "" {
					return true
				}
				return false
			}),
			//CRUD Create(创建)、Read(读取)、Update(更新)和Delete(删除)
			beego.NSNamespace("/blog",
				beego.NSRouter("/", &controllers.BlogController{}),
				// /api/ios/create/topic/
				beego.NSRouter("/:key([0-9]+).html", &controllers.BlogController{}),
			),
			beego.NSRouter("/test", &controllers.TestController{}),
			beego.NSRouter("/matrix", &controllers.MatrixController{}),
		)
	beego.AddNamespace(ns)

}
