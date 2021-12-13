package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/elchead/misou/config"
	"github.com/elchead/misou/search"
	"github.com/urfave/cli/v2"
)

func prettyPrintResults(res []search.SearchResult) {
	for _, s := range res {
		fmt.Printf("Title: %s\nSource: %s\nContent: %s\nProvider: %s\n\n", s.Title, s.Link, s.Content, s.Provider)
	}
}

func main() {
	s := search.NewSearcher()
	config := config.LoadConfig(path.Join(".", "appconfig.json"))
	config.InitSources(s)

	app := &cli.App{
		Name:  "Misou",
		Usage: "explore your knowledge universe",
		Action: func(c *cli.Context) error {
			query := c.Args().Get(0)
			res, err := s.Search(query)
			if err != nil {
				fmt.Printf("Search %s failed: %v\n", query, err)
			}
			prettyPrintResults(res)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
