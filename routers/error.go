package routers

import (
	"github.com/astaxie/beego"
	"html/template"
	"net/http"
)

func Error_init() {
	beego.Errorhandler("404", NotFound)
}

var errtpl = `
<!DOCTYPE html>
<html lang="en">
	<head>
		<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
		<title>{{.Title}}</title>
		<link rel="stylesheet" type="text/css" href="/static/css/error.css" />
	</head>
	<body>
		<div id="wrapper">
			<div id="container">
				<div class="navtop">
					<h1>{{.Title}}</h1>
				</div>
				<div id="content">
					{{.Content}}
					<a href="/" title="Home" class="button">Go Home</a><br />

					<br>Powered by beego {{.BeegoVersion}}
				</div>
			</div>
		</div>
	</body>
</html>
`

// show 404 notfound error.
func NotFound(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("beegoerrortemp").Parse(errtpl)
	// template.New("sd").ParseFiles(...)
	data := make(map[string]interface{})
	data["Title"] = "Page Not Found"
	data["Content"] = template.HTML("<br>Fuck you ." +
		"<br>Perhaps you are here because:" +
		"<br><br><ul>" +
		"<br>The page has moved" +
		"<br>The page no longer exists" +
		"<br>You were looking for your puppy and got lost" +
		"<br>You like 404 pages" +
		"</ul>")
	data["BeegoVersion"] = "1.1..1.1"
	//rw.WriteHeader(http.StatusNotFound)
	t.Execute(rw, data)
}
