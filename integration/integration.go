package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/elchead/misou/config"
	"github.com/elchead/misou/search"
)

var cfgPath string

func SearchUsingIndexer(searcher search.SearcherI,query string) string {
	res, err := searcher.Search(query)
	if err != nil {
		fmt.Printf("Search %s failed: %v\n", query, err)
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(&res)
	return buf.String()
}

func PrepareSearcherWith(indexer search.Indexer) search.SearcherI {
	searcher := search.NewSearcherUsingIndexer(indexer)
	cfg := config.LoadConfig(cfgPath)
	cfg.InitSources(searcher)
	return searcher
}

func PrepareSearcher() search.SearcherI {
	blevePath := filepath.Join(filepath.Dir(cfgPath),"data","misou.bleve")
	return PrepareSearcherWith(search.NewPersistentBleveIndexer(blevePath))
}

func BuildNewIndex() search.SearcherI {
	blevePath := filepath.Join(filepath.Dir(cfgPath),"data","misou.bleve")
	os.RemoveAll(blevePath)
	return PrepareSearcherWith(search.NewPersistentBleveIndexer(blevePath))
}
