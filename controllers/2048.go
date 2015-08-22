package controllers

type G2048Controller struct {
	BaseController
}

//
func (g *G2048Controller) Get() {

	g.AddCustomCssStyle(website.CssUrl+"2048/", "main.css")
	g.AddCustomJsScript(website.JsUrl+"2048/", "bind_polyfill.js", "classlist_polyfill.js", "animframe_polyfill.js",
		"keyboard_input_manager.js", "html_actuator.js", "grid.js", "tile.js",
		"local_storage_manager.js", "game_manager.js", "application.js")

	g.TplNames = "2048.tpl"
}
