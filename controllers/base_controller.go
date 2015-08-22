package controllers

import (
	_ "container/list"

	"github.com/ckeyer/goblog/conf"
	"github.com/ckeyer/goblog/models"
)

var (
	website = conf.GetConf().WebSite
)

type BaseController struct {
	Controller
}

func (this *BaseController) Prepare() {
	this.Ctx.Request.Header.Add("Access-Control-Allow-Origin", "*")

	// 验证是否来自合法域名访问
	if !this.isAllowHost() {
		this.Ctx.WriteString(`<!DOCTYPE html><html><head><meta http-equiv="refresh" content="0; url=` +
			website.HostUrl + string([]byte(this.Ctx.Input.Url())[4:]) + `" /></head></html>`)
		this.StopRun()
	}

	this.Data["Metes"] = ""
	this.Data["Keywords"] = "CKeyer"
	this.Data["Description"] = "CKeyer"
	this.Data["PageTitle"] = "Home"
	this.Data["Styles"] = `<link rel="stylesheet" href="` + website.CssUrl + `style.css" media="screen" type="text/css" />`
	this.Data["Scripts"] = `<script type="text/javascript" src="` + website.JsUrl + `jquery-2.1.3.min.js"></script>
<script type="text/javascript" src="` + website.JsUrl + `default.js"></script>
<script src="` + website.JsUrl + `modernizr.js"></script>
<script src='` + website.JsUrl + `dat.gui.min.js'></script>
<script src='` + website.JsUrl + `toxiclibs.min.js'></script>
<script src='` + website.JsUrl + `animitter.min.js'></script>
<script src="` + website.JsUrl + `bg_index.js"></script>`
	this.Data["CusStyles"] = ``
	this.Data["CusScripts"] = ``
	this.Data["Tail"] = `Download your use my life`

	this.Data["BlogsMonth"] = models.GetBlogsMonth(5)
	this.Data["BlogsTag"] = models.GetHotTags(5)

	this.Layout = "layout/layout.html"
}

// 是否是通过允许的域名访问
func (this *BaseController) isAllowHost() bool {
	for _, v := range website.EnableDomain {
		if this.Ctx.Input.Host() == v {
			return true
		}
	}
	log.Debugf("Not Enable Domain %s\n", this.Ctx.Input.Host())
	return false
}
