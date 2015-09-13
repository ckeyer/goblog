package controllers

import "github.com/astaxie/beego"
import "github.com/ckeyer/goblog/libs"

var (
	log = libs.GetLogger()
)

type Controller struct {
	beego.Controller
}
