package controllers

import (
	_ "container/list"
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/models"
	logpkg "log"
	"os"
)

var (
	log                *logpkg.Logger
	static_url_js_ssl  = " /static/js/"
	static_url_css_ssl = " /static/css/"
	static_url_img_ssl = "/static/img/"
	static_url_js      = " /static/js/"
	static_url_css     = " /static/css/"
	static_url_img     = "/static/img/"
	custom_url_js      = " /static/js/"
	custom_url_css     = " /static/css/"
	custom_url_img     = " /static/img/"
)

func init() {
	log = logpkg.New(os.Stderr, "controller", logpkg.Ltime|logpkg.Lshortfile)
}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {

	this.Data["Title"] = "Download you with my self"
	this.Data["Styles"] = `<link rel="stylesheet" href="` + static_url_css + `style.css" media="screen" type="text/css" />`
	this.Data["Scripts"] = `<script type="text/javascript" src="` + static_url_js + `jquery-2.1.3.min.js"></script>
<script type="text/javascript" src="` + static_url_js + `matrix.js"></script>
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
