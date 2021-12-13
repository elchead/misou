package source

import (
	"os"

	"github.com/elchead/misou/filesystem"
	"github.com/elchead/misou/search"
)

type instapaperEntry struct {
	Content string   `json:"content"`
	Link    string   `json:"href"`
	Title   string `json:"title"`
}

type InstapaperParser struct {
	File *os.File
}

func NewInstapaperParser(f *os.File) *InstapaperParser {
	return &InstapaperParser{f}
}

const instapaperProvider  = "instapaper"

func enrichInstapaperEntries(entries []instapaperEntry) []*search.IndexData {
	res := make([]*search.IndexData, len(entries))
	for i,entry := range entries {
		res[i] = &search.IndexData{Title: entry.Title,Link: entry.Link, Content: entry.Content, Provider: instapaperProvider, ContentType: "url"}
	}
	return res
}

func (i InstapaperParser) TransformToData() ([]*search.IndexData, error) {
	var entries []instapaperEntry
	err := filesystem.LoadVariable(i.File,&entries)
	if err != nil {
		return []*search.IndexData{},err
	}
	return enrichInstapaperEntries(entries),nil
}
