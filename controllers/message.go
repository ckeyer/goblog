package controllers

import (
	"github.com/ckeyer/goblog/lib"
)

type MessageController struct {
	BaseController
}

type InputMessage struct {
	Name    string `form:"name"`
	Email   string `form:"email"`
	Content string `form:"content"`
}

//
func (m *MessageController) Get() {
	m.TplNames = "msgLeave.tpl"
}

func (m *MessageController) Leave() {
	var in InputMessage
	err := m.ParseForm(&in)
	if err != nil {
		log.Error(err.Error())
	}
	if err == nil && lib.IsEmail(in.Email) && len(in.Content) > 0 && len(in.Name) > 3 {
	} else {
		m.Ctx.WriteString(`{"code":"error"}`)
	}
}
