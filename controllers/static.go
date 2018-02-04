package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
)

const (
	staticFolder = "public"
)

// StaticInitializeRoute is used to set the folder routes in the existing router.
// what we want is to make it seem like public/img is really /img.
// so let's walk just the public folder and create individual routes for them.
func StaticInitializeRoute(r chi.Router) {
	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, staticFolder)

	files, err := ioutil.ReadDir(filesDir)

	if err != nil {
		panic("Could not read public directory")
	}

	for _, file := range files {
		if file.IsDir() {
			log.Println("serving folder ", file.Name())
			staticFileServer(r, "/"+file.Name(), http.Dir(filepath.Join(filesDir, file.Name())))
		}
	}
}

func staticFileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}

	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
