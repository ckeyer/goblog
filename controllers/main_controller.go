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
	m.Data["CusStyles"] = `  <link href="/static/css/2048/main.css" rel="stylesheet" type="text/css">
  <link rel="apple-touch-startup-image" href="/static/img/apple-touch-startup-image-640x1096.png" media="(device-width: 320px) and (device-height: 568px) and (-webkit-device-pixel-ratio: 2)"> <!-- iPhone 5+ -->
  <link rel="apple-touch-startup-image" href="/static/img/apple-touch-startup-image-640x920.png"  media="(device-width: 320px) and (device-height: 480px) and (-webkit-device-pixel-ratio: 2)"> <!-- iPhone, retina -->
  <meta name="apple-mobile-web-app-capable" content="yes">
  <meta name="apple-mobile-web-app-status-bar-style" content="black">

  <meta name="HandheldFriendly" content="True">
  <meta name="MobileOptimized" content="320">
  <meta name="viewport" content="width=device-width, target-densitydpi=160dpi, initial-scale=1.0, maximum-scale=1, user-scalable=no, minimal-ui">`

	m.Data["CusScripts"] = `
  <script src="/static/js/2048/bind_polyfill.js"></script>
  <script src="/static/js/2048/classlist_polyfill.js"></script>
  <script src="/static/js/2048/animframe_polyfill.js"></script>
  <script src="/static/js/2048/keyboard_input_manager.js"></script>
  <script src="/static/js/2048/html_actuator.js"></script>
  <script src="/static/js/2048/grid.js"></script>
  <script src="/static/js/2048/tile.js"></script>
  <script src="/static/js/2048/local_storage_manager.js"></script>
  <script src="/static/js/2048/game_manager.js"></script>
  <script src="/static/js/2048/application.js"></script>`

	m.Data["LatestBlogs"] = models.GetBlogs(0, 5)

	m.TplNames = "index.tpl"
}
