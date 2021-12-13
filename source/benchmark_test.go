package source_test

import (
	"testing"

	"github.com/elchead/misou/integration"
	"github.com/elchead/misou/search"
	"github.com/elchead/misou/source"
)
func BenchmarkHistory(b *testing.B) {
	query := "money"
        book := source.NewBrowserHistorySearcher("./history.sh")
        for n := 0; n < b.N; n++ {
                book.Search(query)
        }
}

func BenchmarkBook(b *testing.B) {
	query := "money"
        book := source.NewBookmarkSearcher("./bookmarks.sh")
        for n := 0; n < b.N; n++ {
                book.Search(query)
        }
}

func BenchmarkFile(b *testing.B) {
	query := "money"
        book := source.NewFileSearcher("/Users/adria/homebrew/bin/rga","/Users/adria/Google Drive/Obsidian/Second_brain")
        for n := 0; n < b.N; n++ {
                book.Search(query)
        }
}

var indexer = search.NewBleveIndexer()
var searcher = integration.PrepareSearcherWith(indexer)
func BenchmarkIndexer(b *testing.B) {
        for n := 0; n < b.N; n++ {
                integration.SearchUsingIndexer(searcher,"money")
        }
}
