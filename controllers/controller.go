package controllers

import (
	_ "container/list"
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/lib/logging"
)

var (
	log = logging.GetLogger()
)

type Controller struct {
	beego.Controller
}
