package source

import (
	"testing"

	"github.com/elchead/misou/filesystem"
	"github.com/elchead/misou/search"
	"github.com/stretchr/testify/assert"
)

// func TestReadwiseApi(t *testing.T) {
// 	client := NewReadwiseClient("heREP39IZnUyYdW6dDPoYA8tiflb636X0xXSCYEcgEOAM5B6Wc")
// 	highlights := client.GetData()
// 	res := client.ParseData(highlights)
// 	// assert.NotEmpty(t,res)
// 	AddEntries(res)
// 	// fmt.Printf("%+v",records)
// }

func TestReadCsv(t *testing.T) {
	fname := "../data/readwise-data.csv"
	f, err := filesystem.OpenFile(fname)
	assert.NoError(t, err)
	defer f.Close()
	entries, err := UnmarshalEntriesFromCSV(f)
	assert.NoError(t, err)
	assert.Equal(t,
		"Tweets from Naval", entries[0].Title)
}

func TestReadwiseCsvSearch(t *testing.T) {
	f, err := filesystem.OpenFile("../data/readwise-data.csv")
	assert.NoError(t, err)
	defer f.Close()

	c := ReadwiseCSVParser{f}
	// err = c.Register()
	s := search.NewSearcher()
	s.AddSrcToDb(c)
	searchRes, err := s.Search("money")
	assert.NoError(t, err)
	assert.NotEmpty(t, searchRes)
}
