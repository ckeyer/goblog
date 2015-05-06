package controllers

import (
	_ "container/list"
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/models"
)

type MatrixController struct {
	beego.Controller
}

func (this *MatrixController) Post() {
	wp := models.NewWebPage("首页Matrix")
	wp.IncrViewCount()

	this.Ctx.WriteString(`{"msgcode":1,"data":[[1,2,3,4,0],[1,2,3,4,0],[1,2,3,4,0]]}`)
}
