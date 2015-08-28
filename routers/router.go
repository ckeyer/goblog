package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/ckeyer/goblog/controllers"
	"github.com/ckeyer/goblog/models"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSAfter(log_info),
		beego.NSRouter("/", &controllers.MainController{}),
		beego.NSRouter("/test", &controllers.TestController{}),
		beego.NSRouter("/status", &controllers.StatusController{}, "get:Any;post:Any"),
		beego.NSRouter("/matrix", &controllers.MatrixController{}),
		beego.NSRouter("/note:key([0-9]+).html", &controllers.TagController{}),
		beego.NSRouter("/article_:key([0-9]+).html", &controllers.BlogController{}, "get:ShowBlog"),
		beego.NSNamespace("/blog",
			beego.NSRouter("/", &controllers.BlogController{}, "get:ShowList;post:Post"),
		),
		beego.NSNamespace("/tag",
			beego.NSRouter("/", &controllers.TagController{}),
			beego.NSRouter("/:key([0-9]+).html", &controllers.TagController{}),
		),
		beego.NSNamespace("/msg",
			beego.NSRouter("/", &controllers.MessageController{}),
			beego.NSRouter("/leave", &controllers.MessageController{}, "post:Leave"),
		),
		beego.NSNamespace("/admin",
			beego.NSRouter("/", &controllers.MainController{}),
			beego.NSNamespace("/blog",
				beego.NSRouter("/new", &controllers.BlogController{}, "post:NewBlog;get:AddBlog"),
				beego.NSRouter("/", &controllers.BlogController{}, "get:ShowEditList"),
				beego.NSRouter("/:key([0-9]+).html", &controllers.BlogController{}, "get:EditBlog"),
			),
		),
	)

	// For API
	ns2 := beego.NewNamespace("v2",
		beego.NSBefore(auth),
		beego.NSRouter("/", &controllers.MainController{}),
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(ns2)
}

func log_info(ctx *context.Context) {

	connlog := &models.ConnLog{
		Domain:   ctx.Input.Domain(),
		Host:     ctx.Input.Host(),
		Uri:      ctx.Input.Uri(),
		Ip:       ctx.Input.IP(),
		Scheme:   ctx.Input.Scheme(),
		Method:   ctx.Input.Method(),
		Protocol: ctx.Input.Protocol(),
		Status:   ctx.Output.Status,
	}
	err := connlog.Insert()
	if err != nil {
		println(err.Error())
	}
}

// auth API认证
func auth(ctx *context.Context) {
	sha := ctx.Input.Header("CKEYER_SHA")
	_ = sha
}
