package routers

import (
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/controllers"
	"github.com/ckeyer/goblog/libs"
)

var log = libs.GetLogger()

func init() {
	log.Info("加载路由信息")
	beego.Router("/", &controllers.IndexController{})
}
