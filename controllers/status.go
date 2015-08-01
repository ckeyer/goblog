package controllers

import (
	"github.com/ckeyer/goblog/lib/msg"
)

type StatusController struct {
	Controller
}
type StatusMessage struct {
	msg.Message
	TimeOut int    `form:"timeout"`
	Url     string `form:"url"`
}

// 回显状态信息
func (s *StatusController) Any() {
	s.Data["Url"] = "http://" + s.Ctx.Input.Host()
	s.Data["Msg"] = "error"
	s.Data["TimeOut"] = 0
	m := &StatusMessage{}
	err := s.ParseForm(m)
	if err != nil {
		log.Error(err.Error())
	} else {
		log.Debug(m.Data)
		log.Debug(s.GetString("data"))
		log.Debug(string(m.TimeOut))
		//s.Data["Url"] = m.Url
		s.Data["Msg"] = m.Data
		s.Data["TimeOut"] = m.TimeOut
	}
	s.TplNames = "status.tpl"
}
