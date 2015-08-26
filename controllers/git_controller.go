package controllers

import (
	"github.com/ckeyer/goblog/conf"
	"github.com/ckeyer/goblog/lib/webhook"
)

type WebHookController struct {
	Controller
}

var (
	config = conf.GetConf().WebHook
)

// Post 接受WebHook
func (w *WebHookController) Post() {
	webhook.HookHandle(w.Ctx.ResponseWriter, w.Ctx.Request)
}
