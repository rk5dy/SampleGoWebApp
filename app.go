package main

import (
	"net/http"
	"html/template"
	"io"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func me(res http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(res, "index.gohtml", nil)	
	io.WriteString(res, "My name is Rock Kim")
	HandleError(res, err)
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog")
}

func main() {
	http.HandleFunc("/me", me)
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", nil) 
}

func HandleError(res http.ResponseWriter, err error) {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		io.WriteString(res, "LOL")
	}
}
