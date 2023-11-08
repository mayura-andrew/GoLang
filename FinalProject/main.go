package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/mayura-andrew/GoLang/FinalProject/api"
	"github.com/mayura-andrew/GoLang/FinalProject/data"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from a Go program"))
}
func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("template/index.tmpl")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
	html.Execute(w, data.GetAll())
}
func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", handleHello)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/exhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)
	err := http.ListenAndServe(":3333", server)
	if err == nil {
		fmt.Println("Error while opening the server")
	}

}
