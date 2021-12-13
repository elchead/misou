package search_test

import (
	"testing"

	"github.com/elchead/misou/filesystem"
	"github.com/elchead/misou/search"
	"github.com/stretchr/testify/assert"
)

func TestAddEntry(t *testing.T) {
	txt := "Hi world, I am Adrian and he wants to learn Golang."
	src := search.DummyFile{Content: txt, Title: "Hello"}
	search.AddEntry(src)
	_, err := search.ReadRecord(search.Id("0"))
	assert.NoError(t, err)
}

func TestTranformsearchToData(t *testing.T) {
	txt := "Hi world, I am Adrian and he wants to learn Golang."
	src := search.DummyFile{Content: txt, Title: "Hello"}
	data,_ := src.TransformToData()
	assert.Equal(t, "Hello", data[0].Title)
	assert.Equal(t, txt, data[0].Content)
}

func TestSaveData(t *testing.T) {
	data := search.DummyFile{Title: "Hello", Content: "Hola mundo"}
	search.AddEntry(data)
	file, close := filesystem.CreateTempFile(t, "")
	defer close()
	search.WriteRecordsToFile(file)

	uid := search.Id("0")
	_, err := search.ReadRecord(uid)
	assert.NoError(t, err)
	file.Seek(0, 0)
	_, err = search.ReadRecordFromFile(file, uid)
	assert.NoError(t, err)
}

func TestLoadDb(t *testing.T) {
	search.ResetMemoryDb()
	
	data := search.DummyFile{Title: "Hello", Content: "Hola mundo"}
	search.AddEntry(data)

	recFile, closeR := filesystem.CreateTempFile(t, "")
	defer closeR()
	idxFile, closeI := filesystem.CreateTempFile(t, "")
	defer closeI()
	search.GenerateIndexer()
	err := search.SaveDb(recFile,idxFile)
	assert.NoError(t, err)
	search.ResetMemoryDb()

	recFile.Seek(0,0)
	idxFile.Seek(0,0)
	err = search.LoadDb(recFile,idxFile)
	assert.NoError(t, err)
	
	records,_ := search.Search("mundo")
	assert.Equal(t, 1,len(records))
}
