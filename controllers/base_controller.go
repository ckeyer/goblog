package controllers

import (
	_ "container/list"
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/models"
	logpkg "log"
	"os"
)

var (
	log *logpkg.Logger

	static_url         = beego.AppConfig.String("static_url")
	static_url_js_ssl  = beego.AppConfig.String("static_url_js_ssl")
	static_url_css_ssl = beego.AppConfig.String("static_url_css_ssl")
	static_url_img_ssl = beego.AppConfig.String("static_url_img_ssl")
	static_url_js      = beego.AppConfig.String("static_url_js")
	static_url_css     = beego.AppConfig.String("static_url_css")
	static_url_img     = beego.AppConfig.String("static_url_img")
	custom_url_js      = beego.AppConfig.String("custom_url_js")
	custom_url_css     = beego.AppConfig.String("custom_url_css")
	custom_url_img     = beego.AppConfig.String("custom_url_img")
)

func init() {
	log = logpkg.New(os.Stderr, "controller", logpkg.Ltime|logpkg.Lshortfile)
}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	url_head := this.Ctx.Input.Scheme() + "://" + this.Ctx.Input.Host()
	if url_head+"/" != static_url {
		this.Ctx.WriteString(`<!DOCTYPE html><html><head><meta http-equiv="refresh" content="0; url=` +
			static_url + string([]byte(this.Ctx.Input.Url())[4:]) + `" /></head></html>`)
		this.StopRun()
	}
	this.Data["Metes"] = ""
	this.Data["Keywords"] = "CKeyer"
	this.Data["Description"] = "CKeyer"
	this.Data["PageTitle"] = "Home"
	this.Data["Styles"] = `<link rel="stylesheet" href="` + static_url_css + `style.css" media="screen" type="text/css" />`
	this.Data["Scripts"] = `<script type="text/javascript" src="` + static_url_js + `jquery-2.1.3.min.js"></script>
<script type="text/javascript" src="` + static_url_js + `default.js"></script>
<script src="` + static_url_js + `modernizr.js"></script>
<script src='` + static_url_js + `dat.gui.min.js'></script>
<script src='` + static_url_js + `toxiclibs.min.js'></script>
<script src='` + static_url_js + `animitter.min.js'></script>
<script src="` + static_url_js + `bg_index.js"></script>`
	this.Data["CusStyles"] = ``
	this.Data["CusScripts"] = ``
	this.Data["Tail"] = `Download your use my life`

	this.Data["BlogsMonth"] = models.GetBlogsMonth(5)
	this.Data["BlogsTag"] = models.GetHotTags(5)

	this.Layout = "layout/layout.html"
}
