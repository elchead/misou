package main

import (
	_ "embed"

	"github.com/elchead/misou/integration"
	"github.com/elchead/misou/search"
	"github.com/wailsapp/wails"
)
var searcher search.SearcherI

func basic() string {
	return "World!"
}


func init() {
	searcher = integration.PrepareSearcher()
}
func Search(query string) string {
	return integration.SearchUsingIndexer(searcher,query)
}



//go:embed frontend/build/static/js/main.js
var js string

//go:embed frontend/build/static/css/main.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "misou",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(basic)
	app.Bind(Search)
	app.Run()
}
