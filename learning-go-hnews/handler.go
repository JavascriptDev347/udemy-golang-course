package main

import (
	"fmt"
	"net/http"
)

var htmlContent string = `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>%s</title>
</head>
<body>
%s
</body>
</html>
`

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Home page visited")
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	app.render(w, "index.html", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	aboutContent := `<h1>About</h1>`
	aboutContent = fmt.Sprintf(htmlContent, "About", aboutContent)
	_, _ = w.Write([]byte(aboutContent))
}

func (app *application) contact(writer http.ResponseWriter, request *http.Request) {
	contactContent := `<h1>Contact Page</h1>`
	contactContent = fmt.Sprintf(htmlContent, "Contact", contactContent)
	_, _ = writer.Write([]byte(contactContent))
}
