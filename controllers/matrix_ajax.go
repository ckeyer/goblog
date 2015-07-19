package controllers

import (
	_ "container/list"
	// "encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/ckeyer/goblog/models"
	"strconv"
)

type MatrixController struct {
	beego.Controller
}

func (this *MatrixController) Post() {
	// log.Println("is https: ", this.Ctx.Input.IsSecure())
	wp := models.NewWebPage("首页Matrix")
	wp.IncrViewCount()

	resmag := &models.MatrixUpJson{}
	if err := this.ParseForm(resmag); err != nil {
		this.Ctx.WriteString(`{"code":-1,"data":"` + err.Error() + `"}`)
		return
	}

	switch resmag.Code {
	case 1:
		vals, err := models.GetAllMatrix()
		if err != nil {
			this.Ctx.WriteString(`{"code":-2,"data":"models.GetAllMatrix"}`)
			return
		}
		jsonstr := vals.ToJson()
		this.Ctx.WriteString(`{"code":1,"data":` + jsonstr + `}`)
	case 2:
		resmag.H, _ = strconv.Atoi(this.GetString("h"))
		resmag.W, _ = strconv.Atoi(this.GetString("w"))
		resmag.Col, _ = strconv.Atoi(this.GetString("val"))

		b, e := models.UpdateMatrix(resmag.H, resmag.W)
		if e != nil {
			log.Error(e.Error())
			this.Ctx.WriteString(`{"code":-3,"data":"up error"}`)
		} else if b {
			this.Ctx.WriteString(`{"code":0,"data":"up false"}`)
		} else {
			this.Ctx.WriteString(`{"code":0,"data":"up success"}`)
		}
	default:
		this.Ctx.WriteString(`{"code":-1,"data":"none"}`)
	}
}
