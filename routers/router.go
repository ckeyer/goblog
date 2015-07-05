package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/ckeyer/goblog/controllers"
	"github.com/ckeyer/goblog/models"
)

func init() {
	controllers.Test()
	ns := beego.NewNamespace("/v1",
		beego.NSBefore(resource_init),
		beego.NSAfter(log_info),
		beego.NSRouter("/", &controllers.MainController{}),
		beego.NSRouter("/test", &controllers.TestController{}),
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
		beego.NSNamespace("/admin",
			beego.NSRouter("/", &controllers.MainController{}),
			beego.NSNamespace("/blog",
				beego.NSRouter("/new", &controllers.BlogController{}, "post:NewBlog;get:AddBlog"),
				beego.NSRouter("/", &controllers.BlogController{}, "get:ShowEditList"),
				beego.NSRouter("/:key([0-9]+).html", &controllers.BlogController{}, "get:EditBlog"),
			),
		),
		beego.NSRouter("/chat", &controllers.ChatController{}),
		beego.NSRouter("/test", &controllers.TestController{}),
	)
	beego.AddNamespace(ns)
}

func resource_init(ctx *context.Context) {
	if ctx.Input.IsSecure() {
		ctx.Output.Session("STATIC_URL_JS", beego.AppConfig.String("static_url_js_ssl"))
		ctx.Output.Session("STATIC_URL_CSS", beego.AppConfig.String("static_url_css_ssl"))
		ctx.Output.Session("STATIC_URL_IMG", beego.AppConfig.String("static_url_img_ssl"))
	} else {
		ctx.Output.Session("STATIC_URL_JS", beego.AppConfig.String("static_url_js"))
		ctx.Output.Session("STATIC_URL_CSS", beego.AppConfig.String("static_url_css"))
		ctx.Output.Session("STATIC_URL_IMG", beego.AppConfig.String("static_url_img"))
	}
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
