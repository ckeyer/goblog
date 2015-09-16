package controllers

import (
	_ "container/list"

	"github.com/ckeyer/goblog/conf"
)

type BaseController struct {
	Controller
}

func (b *BaseController) Prepare() {
	website = conf.GetConf().WebSite

	b.Ctx.Request.Header.Add("Access-Control-Allow-Origin", "*")

	// 验证是否来自合法域名访问
	if !b.IsAllowHost() {
		b.Ctx.WriteString(`<!DOCTYPE html><html><head><meta http-equiv="refresh" content="0; url=` +
			website.HostUrl + string([]byte(b.Ctx.Input.Url())[4:]) + `" /></head></html>`)
		b.StopRun()
	}

	b.Data["WebSite"] = website
	b.Data["HostUrl"] = website.HostUrl
	b.SetPageTitle(website.Title)
	b.SetDescript(website.Description)
	b.AddKeyWord(website.Keywords...)

	b.Data["Metes"] = ""
	b.AddCustomCssStyle("//cdn.bootcss.com/bootstrap/3.3.5/css/", "bootstrap.min.css", "bootstrap-theme.min.css")
	b.AddCustomCssStyle("//cdn.bootcss.com/font-awesome/4.4.0/css/", "font-awesome.min.css")
	b.AddCssStyle("style.css")
	//	b.AddCustomCssStyle("http://fonts.useso.com/", "css?family=Open+Sans:300,400,600&subset=latin,latin-ext")

	b.AddCustomJsScript("//cdn.bootcss.com/jquery/2.1.4/", "jquery.min.js")
	b.AddCustomJsScript("//cdn.bootcss.com/bootstrap/3.3.5/css/", "bootstrap.min.js")
	b.AddCustomCssStyle("//cdn.bootcss.com/jquery-migrate/1.2.1/", "jquery-migrate.min.js")
	b.AddCustomJsScript("//cdn.bootcss.com/wow/1.1.2/", "wow.min.js")

	b.Data["Tail"] = `Download your use my life`

	b.LayoutSections = make(map[string]string)

	b.Layout = "layout/index.html"
}
