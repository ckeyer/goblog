package controllers

import (
	_ "container/list"

	"fmt"
	"strings"

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

	this.Data["WebSite"] = website
	this.SetPateTitle("Home")
	this.SetDescript("Ckeyer blog")
	this.AddKeyWord("ckeyer")
	this.Data["Metes"] = ""
	this.AddCssStyle("style.css")
	this.AddJsScript("jquery-2.1.3.min.js", "default.js", "modernizr.js", "dat.gui.min.js", "toxiclibs.min.js", "animitter.min.js", "bg_index.js")
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

//  (this *BaseController)SetPateTitle 设置页面显示标题
func (this *BaseController) SetPateTitle(title string) {
	this.Data["PageTitle"] = title
}

// (this *BaseController)SetKeyWord 设置或更新Meta关键字
func (this *BaseController) SetKeyWord(args ...string) {
	this.Data["Keywords"] = strings.Join(args, ",")
}

// (this *BaseController)AddKeyWord 添加Meta关键字
func (this *BaseController) AddKeyWord(args ...string) {
	if this.Data["Keywords"] != nil {
		keyword := strings.Split(fmt.Sprint(this.Data["Keywords"]), ",")
		this.Data["Keywords"] = strings.Join(append(keyword, args...), ",")
	} else {
		this.SetKeyWord(args...)
	}
}

// (this *BaseController)SetDescript 设置Mets描述
func (this *BaseController) SetDescript(des string) {
	this.Data["Description"] = des
}

// (this *BaseController)AddJsScript 按配置的js路径添加js文件
func (this *BaseController) AddJsScript(args ...string) {
	this.AddCustomJsScript(website.JsUrl, args...)
}

// (this *BaseController)AddCustomJsScript 添加自定义js
func (this *BaseController) AddCustomJsScript(src_url string, args ...string) {
	var jstags []string

	if this.Data["Scripts"] != nil {
		jstags = strings.Split(fmt.Sprint(this.Data["Scripts"]), "\n")
	}
	for _, js := range args {
		newtag := fmt.Sprintf(`<script type="text/javascript" src="%s%s"></script>`, src_url, js)
		jstags = append(jstags, newtag)
	}
	this.Data["Scripts"] = strings.Join(jstags, "\n")

}

// (this *BaseController)AddCssStyle 按配置的css路径添加css文件
func (this *BaseController) AddCssStyle(args ...string) {
	this.AddCustomCssStyle(website.CssUrl, args...)
}

// (this *BaseController)AddCustomCssStyle 添加自定义css
func (this *BaseController) AddCustomCssStyle(src_url string, args ...string) {
	var csstags []string

	if this.Data["Styles"] != nil {
		csstags = strings.Split(fmt.Sprint(this.Data["Styles"]), "\n")
	}
	for _, css := range args {
		newtag := fmt.Sprintf(`<link rel="stylesheet" media="screen" type="text/css"  href="%s%s"/>`, src_url, css)
		csstags = append(csstags, newtag)
	}
	this.Data["Styles"] = strings.Join(csstags, "\n")
}
