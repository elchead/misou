package source

import (
	"testing"

	"github.com/elchead/misou/filesystem"
	"github.com/elchead/misou/search"
	"github.com/stretchr/testify/assert"
)

func TestInstapaper(t *testing.T){
	f, err := filesystem.OpenFile("../data/instapaper-full-text.json")
	assert.NoError(t, err)
	defer f.Close()

	c := InstapaperParser{f}
	s := search.NewSearcher()
	s.AddSrcToDb(c)
	searchRes, err := s.Search("tdd")
	assert.NoError(t, err)
	assert.NotEmpty(t, searchRes)

}
