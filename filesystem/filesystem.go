package filesystem

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func CreateTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "*db.json")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}

func createFile(path string) (*os.File, error) {
	f, errCreating := os.Create(path)
	if errCreating != nil {
		log.Fatal("Error, could not create database for path: ", path, " with: ", errCreating)
	}
	return f, errCreating
}

func OpenFile(path string) (*os.File, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		jsonFile, err = createFile(path)
	}
	return jsonFile, err
}

func LoadVariable(file io.Reader, v interface{}) error {
	return jsoniter.NewDecoder(file).Decode(&v)
}

func WriteVariable(file io.Writer, v interface{}) error {
	return jsoniter.NewEncoder(file).Encode(&v)
}
