package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// RootRouter is the chi router to mount  set in Init
var RootRouter *chi.Mux

var templates = template.Must(template.ParseFiles(
	"./views/root.tmpl"))

func init() {
	RootRouter = chi.NewRouter()
	RootRouter.Get("/", homeHandler)
	RootRouter.Get("/siteicon.ico", faviconHandler)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "root.tmpl", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/img/favicon.ico")
}
