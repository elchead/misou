package search_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/elchead/misou/search"
	"github.com/stretchr/testify/assert"
)

func getField(v *search.SearchResult, field string) reflect.Value {
	r := reflect.ValueOf(v)
	return reflect.Indirect(r).FieldByName(field)
    }

func MapContent(vs []search.SearchResult) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
	    vsm[i] = v.Content
	}
	return vsm
    }

func MapContentString(vs []search.SearchResult,field string) []string {
vsm := make([]string, len(vs))
for i, v := range vs {
	vsm[i] = getField(&v,field).String()
}
return vsm
}

func MapContentFloat(vs []search.SearchResult,field string) []float64 {
	vsm := make([]float64, len(vs))
	for i, v := range vs {
		vsm[i] = getField(&v,field).Float()
	}
	return vsm
	}

func TestBlevel(t *testing.T) {
	sut := search.NewBleveIndexer()
	sut.AddEntry(&search.IndexData{Title:"Hello world",Content:"nothing is fragile"})
	sut.AddEntry(&search.IndexData{Title:"Hello",Content:"is it antifragile. or dat fragile"})
	res, err := sut.Search("fragile")
	assert.NoError(t, err)
	assert.Equal(t,2, len(res))

	contents := MapContent(res)
	assert.Contains(t,contents,"nothing is fragile")
}

func TestBlevelOpen(t *testing.T) {
	// search.OpenBleveIndexer()

}



func TestBleveSearch(t *testing.T) {
	sut := search.NewBleveIndexer()
	sut.AddEntry(&search.IndexData{Title:"Hello world",Content:"i have been to south america"})	
	sut.AddEntry(&search.IndexData{Title:"Hello south",Content:"south africa is great"})	
	sut.AddEntry(&search.IndexData{Title:"Hello south",Content:"go southern"})
	t.Run("exact search", func(t *testing.T) {
		res, err := sut.Search("\"South America\"")
		assert.NoError(t, err)
		assert.Equal(t,1, len(res))
	})

	t.Run("match search", func(t *testing.T) {
		res, err := sut.Search("South America")
		assert.NoError(t, err)
		assert.Equal(t,2, len(res))
	
		contents := MapContent(res)
		assert.Contains(t,contents,"i have been to south america")	
	})
}
func TestMultipleResults(t *testing.T) {
	t.Run("batch if same provider", func(t *testing.T) {
			sut := search.NewBleveIndexer()
			sut.AddEntry(&search.IndexData{Title:"Hello",Content:"go south",Provider:"google"})
			sut.AddEntry(&search.IndexData{Title:"Hello world",Content:"i have been to south america",Provider:"instapaper"})	
			sut.AddEntry(&search.IndexData{Title:"Hello world",Content:"south africa is great",Provider:"instapaper"})	
			res, err := sut.Search("south")
			assert.NoError(t, err)
			assert.Equal(t,2, len(res))
		
			contents := MapContent(res)
			assert.NotContains(t,contents,"i have been to south america")	
	})
	t.Run("don't batch if different provider", func(t *testing.T) {
		sut := search.NewBleveIndexer()
		sut.AddEntry(&search.IndexData{Title:"Hello",Content:"go south",Provider:"google"})
		sut.AddEntry(&search.IndexData{Title:"Hello world",Content:"i have been to south america",Provider:"readwise"})	
		sut.AddEntry(&search.IndexData{Title:"Hello world",Content:"south africa is great",Provider:"instapaper"})	
		res, err := sut.Search("south")
		assert.NoError(t, err)
		fmt.Println(res)
		assert.Equal(t,3, len(res))
})	
}

func TestScoresIncluded(t *testing.T) {
	sut := search.NewBleveIndexer()
	sut.AddEntry(&search.IndexData{Title:"Hello",Content:"go south",Provider:"google"})
	sut.AddEntry(&search.IndexData{Title:"Hello world",Content:"i have been to south america",Provider:"instapaper"})	
	sut.AddEntry(&search.IndexData{Title:"Hello world",Content:"south africa is great",Provider:"instapaper"})	
	res, _ := sut.Search("south")
	scores := MapContentFloat(res,"Score")
	assert.NotContains(t,scores,0.)
}
