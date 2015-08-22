package controllers

import (
	_ "container/list"

	"github.com/ckeyer/goblog/models"
)

type MainController struct {
	BaseController
}

func (m *MainController) Get() {
	m.AddKeyWord("极客", "个人博客")

	m.AddCustomCssStyle(website.CssUrl+"2048/", "main.css")
	m.AddCustomJsScript(website.JsUrl+"2048/", "bind_polyfill.js", "classlist_polyfill.js", "animframe_polyfill.js",
		"keyboard_input_manager.js", "html_actuator.js", "grid.js", "tile.js",
		"local_storage_manager.js", "game_manager.js", "application.js")

	m.Data["LatestBlogs"] = models.GetBlogs(0, 5)

	m.TplNames = "index.tpl"
}
