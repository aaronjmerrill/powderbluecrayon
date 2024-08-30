package main

import (
	"log"
	"net/http"

	"powderbluecrayon/templates"
	"powderbluecrayon/templates/layouts"
	gameLayouts "github.com/powderbluecrayon/memory-game/templates/layouts"

	"github.com/a-h/templ"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.Handle("/", templ.Handler(templates.Root(layouts.Resume())))
	http.Handle("/game", templ.Handler(templates.Root(gameLayouts.GameLayout())))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
