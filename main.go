package main

import (
	"log"
	"net/http"

	"powderbluecrayon/templates"
	"powderbluecrayon/templates/layouts"

	"github.com/a-h/templ"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", templ.Handler(templates.Root(layouts.Index())))
	http.Handle("/calculators", templ.Handler(templates.Root(layouts.Index())))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
