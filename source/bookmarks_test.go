package source_test

import (
	"testing"

	"github.com/elchead/misou/config"
	"github.com/elchead/misou/source"
	"github.com/stretchr/testify/assert"
)

func TestBookmarks(t *testing.T) {
	query := "startup"
	cfg := config.LoadConfig(source.AppConfigPath)
	sut := source.NewBookmarkSearcher(cfg.BookmarkPath())
	res, err := sut.Search(query)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.NotEmpty(t, res[0].Title)
	assert.NotEmpty(t, res[0].Link)
	// assert.Equal(t, true, res)
}

func TestNoBookmarksFound(t *testing.T) {
	query := "supercalifragilistic"
	config := config.LoadConfig(source.AppConfigPath)
	sut := source.NewBookmarkSearcher(config.BookmarkPath())
	res, err := sut.Search(query)
	assert.NoError(t, err)
	assert.Empty(t, res)
}
