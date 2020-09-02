package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var t *template.Template
var funcMap *template.FuncMap

func init() {
	t = template.New("date_format")
	funcMap = &template.FuncMap{
		"dateFormat": dateFormat,
	}
	t.Funcs(*funcMap)

	t = template.Must(template.New("date_format.html").Funcs(*funcMap).ParseFiles("templates/date_format.html", "templates/head.html"))
}

func main() {
	http.HandleFunc("/", serveTemplate)
	http.ListenAndServe(":9090", nil)
}

func serveTemplate(rw http.ResponseWriter, req *http.Request) {
	data := struct {
		Title string
		Date  time.Time
	}{
		Title: "Date Format",
		Date:  time.Now(),
	}
	fmt.Printf("Serving template\n")
	t.Execute(rw, data)
}

func dateFormat(layout string, d time.Time) string {
	return d.Format(layout)
}
