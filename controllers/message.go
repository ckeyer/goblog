package controllers

import (
	_ "encoding/json"
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
	log.Debug(m.GetString("content"))
	log.Debug("FUCK")
	log.Debug(in.Name)
	log.Debug(in.Content)
	m.Ctx.WriteString(`<!DOCTYPE html><html><head><meta http-equiv="refresh" content="333; url=` +
		STATIC_URL + `" /></head><body>感谢您反馈...</body></html>`)
}
