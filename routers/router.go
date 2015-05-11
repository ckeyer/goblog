package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/ckeyer/goblog/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/test", &controllers.TestController{}),
		beego.NSRouter("/", &controllers.MainController{}),
		beego.NSRouter("/matrix", &controllers.MatrixController{}),
		beego.NSNamespace("/blog",
			beego.NSRouter("/", &controllers.BlogController{}),
			beego.NSRouter("/:key([0-9]+).html", &controllers.BlogController{}),
		),
		beego.NSNamespace("/tag",
			beego.NSRouter("/", &controllers.TagController{}),
			beego.NSRouter("/:key([0-9]+).html", &controllers.TagController{}),
		),
		beego.NSRouter("/note:key([0-9]+).html", &controllers.TagController{}),
		beego.NSNamespace("/admin",
			beego.NSCond(func(ctx *context.Context) bool {
				if ctx.Input.IsSecure() {
					return true
				}
				return false
			}),
			beego.NSRouter("/", &controllers.MainController{}),
		),
	)
	beego.AddNamespace(ns)
}
