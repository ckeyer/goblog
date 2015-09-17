package controllers

import (
	"github.com/ckeyer/goblog/models"
)

type WebhookController struct {
	Controller
}

// (w *WebhookController)Post ...
func (w *WebhookController) Post() {
	req := w.Ctx.Request
	res := w.Ctx.ResponseWriter
	models.DoWebhook(res, req)
	w.Ctx.WriteString("")
	return
}
