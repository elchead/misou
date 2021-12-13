package source

import (
	"testing"

	"github.com/elchead/misou/filesystem"
	"github.com/elchead/misou/search"
	"github.com/stretchr/testify/assert"
)

func TestTwitter(t *testing.T) {
	f, err := filesystem.OpenFile("../data/tweet.json")
	assert.NoError(t, err)
	defer f.Close()
	sut := TwitterParser{f}
	s := search.NewSearcher()
	s.AddSrcToDb(sut)
	searchRes, err := s.Search("website")
	assert.NoError(t, err)
	assert.NotEmpty(t, searchRes)
}
