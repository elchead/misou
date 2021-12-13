package filesystem

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func writeToFile(t testing.TB,name string){
	f,err := OpenFile(name)
	assert.NoError(t, err)
	defer f.Close()	
	f.WriteString("Hi")
}

func TestOpenFileTwice(t *testing.T) {
	fname := "text.txt"
	writeToFile(t,fname)
	writeToFile(t,fname)
	defer os.Remove(fname)
}

type dummy struct {
	Name string `json:"name"`
	Num int `json:"num"`
}

func TestWriteAndLoadVariable(t *testing.T) {
	f,close := CreateTempFile(t,"")
	defer close()
	dummyVar := dummy{"dummy",1}
	err := WriteVariable(f,dummyVar)
	assert.NoError(t, err)

	f.Seek(0,0)
	var loadDummy dummy
	err = LoadVariable(f,&loadDummy)
	assert.NoError(t, err)
	assert.Equal(t,dummyVar,loadDummy)
}
