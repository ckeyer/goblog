package controllers

type ChatController struct {
	BaseController
}

func (this *ChatController) Get() {

	this.TplNames = "chat.tpl"
}
