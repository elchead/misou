package source_test

import (
	"testing"

	"github.com/elchead/misou/config"
	"github.com/elchead/misou/search"
	"github.com/elchead/misou/source"
	"github.com/stretchr/testify/assert"
)

func foundDuplicate(res []search.SearchResult) bool {
	visited := make(map[string]bool, 0)
	for _, r := range res {
		if visited[r.Title] == true {
			return true
		} else {
			visited[r.Title] = true
		}
	}
	return false
}

func TestFoundDuplicate(t *testing.T) {
	t.Run("find duplicates", func(t *testing.T) {
		data := []search.SearchResult{{Title: "A"}, {Title: "B"}, {Title: "A"}}
		assert.Equal(t, true, foundDuplicate(data))
	})
	t.Run("find no duplicate", func(t *testing.T) {
		data := []search.SearchResult{{Title: "A"}, {Title: "B"}, {Title: "AA"}}
		assert.Equal(t, false, foundDuplicate(data))
	})
}
// TODO why fragile? (need to run h command inside my shell first?)
func TestHistory(t *testing.T) {
	query := "money"
	cfg := config.LoadConfig(source.AppConfigPath)
	sut := source.NewBrowserHistorySearcher(cfg.HistoryPath())
	res, err := sut.Search(query)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res[0].Title)
	// assert.NotEmpty(t, res[0].Link)
}

func TestNoHistoryFound(t *testing.T) {
	query := "supercalifragilisticoooo"
	cfg := config.LoadConfig(source.AppConfigPath)
	sut := source.NewBrowserHistorySearcher(cfg.HistoryPath())
	res, err := sut.Search(query)
	assert.NoError(t, err)
	assert.Empty(t, res)
}

func TestNoDuplicates(t *testing.T) {
	query := "money"
	cfg := config.LoadConfig(source.AppConfigPath)
	sut := source.NewBrowserHistorySearcher(cfg.HistoryPath())
	res, err := sut.Search(query)
	assert.NoError(t, err)
	assert.Equal(t, false, foundDuplicate(res))
}
