package integration

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/elchead/misou/search"
	"github.com/stretchr/testify/assert"
)
func assertApiProviderResultsExist(t testing.TB,res []search.SearchResult){
	providers := make(map[string]bool)
	for _, e := range res {
		providers[e.Provider] = true
	}
	assert.Contains(t,providers,"file")
}

func assertResultsAreOrdered(t testing.TB,res []search.SearchResult){
	t.Helper()
	for i := 0; i < len(res)-1; i++ {
		assert.True(t,res[i].Score >= res[i+1].Score)
	}
}

func assertSearchUsingIndexer(t testing.TB,indexer search.Indexer){
	searcher := PrepareSearcherWith(indexer)

	res := SearchUsingIndexer(searcher,"tdd")
	
	var searchRes []search.SearchResult
	if err := json.Unmarshal([]byte(res), &searchRes); err != nil {
		log.Fatal(err)
	}
	assert.NotEmpty(t, searchRes)
	assertApiProviderResultsExist(t,searchRes)
	assertResultsAreOrdered(t,searchRes)

	res = SearchUsingIndexer(searcher,"journal")
	var searchRes2 []search.SearchResult
	json.Unmarshal([]byte(res), &searchRes2)
	assert.NotEmpty(t, searchRes2)	
}
func TestIndexers(t *testing.T) {
	t.Run("apollo indexer",func(t *testing.T){
		assertSearchUsingIndexer(t,search.NewApolloIndexer())
	})
	t.Run("bleve indexer",func(t *testing.T){
		assertSearchUsingIndexer(t,search.NewBleveIndexer())
		
	})
}

func TestFillwordSearchIsEmpty(t *testing.T) {
	searcher := PrepareSearcherWith(search.NewBleveIndexer())
	res, err := searcher.Search("to")
	assert.NoError(t, err)
	assert.Empty(t,res)
}
