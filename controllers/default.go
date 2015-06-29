package controllers

import (
	_ "container/list"
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/models"
	logpkg "log"
	"os"
)

var log *logpkg.Logger

func init() {
	log = logpkg.New(os.Stderr, "controller", logpkg.Ltime|logpkg.Lshortfile)
}

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Prepare() {

	this.Data["Title"] = "Download you with my self"
	this.Data["Scripts"] = `<link rel="stylesheet" href="/static/css/style.css" media="screen" type="text/css" />`
	this.Data["HtmlHead"] = `<script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
<script type="text/javascript" src="/static/js/matrix.js"></script>
<script src="/static/js/modernizr.js"></script>
<script src='/static/js/dat.gui.min.js'></script>
<script src='/static/js/toxiclibs.min.js'></script>
<script src='/static/js/animitter.min.js'></script>
<script src="/static/js/bg_index.js"></script>`

	this.Data["BlogsMonth"] = models.GetBlogsMonth(5)
	this.Data["BlogsTag"] = models.GetHotTags(5)

	this.Layout = "layout/layout.html"
}
