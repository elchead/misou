package main

import (
	"log"
	"net/http"
	"path"

	"github.com/elchead/misou/config"
	"github.com/elchead/misou/search"
	"github.com/elchead/misou/webserver"
)

func main() {
	searcher := search.NewSearcher()
	config := config.LoadConfig(path.Join("..", "appconfig.json"))
	config.InitSources(searcher)
	server := webserver.NewSearchHandler(searcher)
	log.Fatal(http.ListenAndServe("localhost:5000", server))
}
