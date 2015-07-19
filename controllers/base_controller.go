package controllers

import (
	_ "container/list"
	"github.com/astaxie/beego"
	logpkg "github.com/ckeyer/go-log"
	"github.com/ckeyer/goblog/models"
	"os"
)

var (
	logfmt = logpkg.MustStringFormatter(
		"%{time:15:04:05} %{shortfile} %{longfunc}%{color} ▶ %{color:reset}[%{color}%{level:.4s}%{color:reset}] %{message}")
	log                = logpkg.MustGetLogger("example")
	STATIC_URL         = beego.AppConfig.String("static_url")
	STATIC_URL_JS_SSL  = beego.AppConfig.String("static_url_js_ssl")
	STATIC_URL_CSS_SSL = beego.AppConfig.String("static_url_css_ssl")
	STATIC_URL_IMG_SSL = beego.AppConfig.String("static_url_img_ssl")
	STATIC_URL_JS      = beego.AppConfig.String("static_url_js")
	STATIC_URL_CSS     = beego.AppConfig.String("static_url_css")
	STATIC_URL_IMG     = beego.AppConfig.String("static_url_img")
	custom_url_js      = beego.AppConfig.String("custom_url_js")
	custom_url_css     = beego.AppConfig.String("custom_url_css")
	custom_url_img     = beego.AppConfig.String("custom_url_img")

	ALLOW_HOSTS = []string{"http://localhost/", "http://ingdown.com/", "http://www.ckeyer.com/"}
)

func init() {
	backend1 := logpkg.NewLogBackend(os.Stderr, "", 0)
	backend2 := logpkg.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backend2Formatter := logpkg.NewBackendFormatter(backend2, logfmt)

	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logpkg.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logpkg.ERROR, "")

	// Set the backends to be used.
	logpkg.SetBackend(backend1Leveled, backend2Formatter)

	// log = logpkg.New(os.Stderr, "controller", logpkg.Ltime|logpkg.Lshortfile)
	// log = logpkg.new
}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {
	if !this.isAllowHost() {
		this.Ctx.WriteString(`<!DOCTYPE html><html><head><meta http-equiv="refresh" content="0; url=` +
			STATIC_URL + string([]byte(this.Ctx.Input.Url())[4:]) + `" /></head></html>`)
		this.StopRun()
	}
	this.Data["Metes"] = ""
	this.Data["Keywords"] = "CKeyer"
	this.Data["Description"] = "CKeyer"
	this.Data["PageTitle"] = "Home"
	this.Data["Styles"] = `<link rel="stylesheet" href="` + STATIC_URL_CSS + `style.css" media="screen" type="text/css" />`
	this.Data["Scripts"] = `<script type="text/javascript" src="` + STATIC_URL_JS + `jquery-2.1.3.min.js"></script>
<script type="text/javascript" src="` + STATIC_URL_JS + `default.js"></script>
<script src="` + STATIC_URL_JS + `modernizr.js"></script>
<script src='` + STATIC_URL_JS + `dat.gui.min.js'></script>
<script src='` + STATIC_URL_JS + `toxiclibs.min.js'></script>
<script src='` + STATIC_URL_JS + `animitter.min.js'></script>
<script src="` + STATIC_URL_JS + `bg_index.js"></script>`
	this.Data["CusStyles"] = ``
	this.Data["CusScripts"] = ``
	this.Data["Tail"] = `Download your use my life`

	this.Data["BlogsMonth"] = models.GetBlogsMonth(5)
	this.Data["BlogsTag"] = models.GetHotTags(5)

	this.Layout = "layout/layout.html"
}

// 是否是通过允许的域名访问
func (this *BaseController) isAllowHost() bool {
	url_head := this.Ctx.Input.Scheme() + "://" + this.Ctx.Input.Host()
	if url_head+"/" == STATIC_URL {
		return true
	}
	for _, v := range ALLOW_HOSTS {
		if url_head+"/" == v {
			return true
		}
	}
	return false
}
