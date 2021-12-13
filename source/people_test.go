package source

import (
	"testing"

	"github.com/elchead/misou/filesystem"
	"github.com/elchead/misou/search"
	"github.com/stretchr/testify/assert"
)

func TestPeople(t *testing.T){
	f, err := filesystem.OpenFile("../data/people.json")
	assert.NoError(t, err)
	defer f.Close()

	c := NewPeopleParser(f)
	s := search.NewSearcher()
	s.AddSrcToDb(c)
	searchRes, err := s.Search("Canada")
	assert.NoError(t, err)
	assert.NotEmpty(t, searchRes)
}
