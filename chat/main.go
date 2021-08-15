package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"sync"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	t.templ.Execute(w, nil)
}

func main() {
	templateFile := "chat.html"
	httpServerDescription := ":8080"
	r := newRoom()

	log.Printf("Serving the template file %s at %s\n", templateFile, httpServerDescription)

	http.Handle("/", &templateHandler{filename: templateFile})
	http.Handle("/room", r)

	go r.run()

	if err := http.ListenAndServe(httpServerDescription, nil); err != nil {
		log.Fatal("ListenAnServe:", err)
	}
}
