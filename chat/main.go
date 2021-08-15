package main

import (
	"flag"
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

	t.templ.Execute(w, r)
}

func main() {
	var templateFile = flag.String("template", "chat.html", "The template file for the chat page.")
	var addr = flag.String("addr", ":8080", "The addt of the application.")
	flag.Parse()

	r := newRoom()

	log.Printf("Serving the template \"%s\" at %s\n", *templateFile, *addr)

	http.Handle("/", &templateHandler{filename: *templateFile})
	http.Handle("/room", r)

	go r.run()

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAnServe:", err)
	}
}
