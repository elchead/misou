package source

import (
	"fmt"
	"os"

	"github.com/elchead/misou/search"
	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
)

var readwiseProvider = "readwise"

// var baseUrl = "https://readwise.io/api/v2/"
// type ReadwiseClient struct {
// 	http *http.Client
// 	token string
// }

// func NewReadwiseClient(token string) *ReadwiseClient {
// 	return &ReadwiseClient{&http.Client{},"Token "+token}
// }

// type payload struct {
// 	Count int `json:"count"`
// 	Results []result `json:"results"`
// }

// type result struct {
// 	Id int `json:"id"`
// 	Text string `json:"text"`
// 	Note    string `json:"note"`
// 	HiglightedAt time.Time `json:"highlighted_at"`
// 	Url string `json:"url"`
// 	BookId int `json:"book_id"`
// 	Tags []string `json:"tags"`
// }

// //only returns small subset and incomplete info about title, etc from HTTP request.
// func (c *ReadwiseClient) GetData() *payload {
// 	req, _ := http.NewRequest("GET",baseUrl+"highlights/",nil)
// 	req.Header.Set("Authorization",c.token)
// 	resp, _ := c.http.Do(req)
// 	var data *payload
// 	jsoniter.NewDecoder(resp.Body).Decode(&data)
// 	return data
// }

// func (c *ReadwiseClient) ParseData(data *payload) []*search.IndexData {
// 	fmt.Printf("Len: %d, Count: %d",len(data.Results),data.Count)
// 	res := make([]*search.IndexData, len(data.Results))
// 	for i,entry := range data.Results {
// 		res[i] = &search.IndexData{Title: "",Link: entry.Url, Content: fmt.Sprintf("%s\nNote: %s", entry.Text,entry.Note), Provider: readwiseProvider, ContentType: "text"}
// 	}
// 	return res
// }


type readwiseEntry struct {
	Highlight string `csv:"Highlight"`
	Title string `csv:"Book Title"`
	Author string `csv:"Book Author"`
	Note string `csv:"Note"`
}


func UnmarshalEntriesFromCSV(in *os.File) ([]*readwiseEntry,error) {
	var entries []*readwiseEntry
	err := gocsv.UnmarshalFile(in, &entries)
	return entries, err
}


type ReadwiseCSVParser struct {
	File *os.File
}

func NewReadwiseCSVParser(file *os.File) *ReadwiseCSVParser {
	c := &ReadwiseCSVParser{File: file}
	return c
}

// TODO refactor to care less about search implementation
func (c ReadwiseCSVParser) TransformToData() ([]*search.IndexData,error) {
	entries, err := UnmarshalEntriesFromCSV(c.File)
	if err != nil {
		return nil,errors.Wrap(err,"unable to unmarshal entries from Readwise CSV")
	}		
	res := make([]*search.IndexData, len(entries))
	for i,entry := range entries {
		res[i] = &search.IndexData{Title: entry.Title,Link: "", Content: fmt.Sprintf("%s\nNote: %s", entry.Highlight,entry.Note), Provider: readwiseProvider, ContentType: "text"}
	}
	return res,nil
}

func (c *ReadwiseCSVParser) Search(keywords string) ([]search.SearchResult,error) {
	return search.Search(keywords)
}

// TODO delete
// func (r ReadwiseCSVParser) Register() error {
// 	entries, err := UnmarshalEntriesFromCSV(r.File)
// 	if err != nil {
// 		return errors.Wrap(err,"unable to unmarshal entries from Readwise CSV")
// 	}
// 	data := r.TransformToData(entries)	
// 	AddEntries(data)
// 	GenerateIndexer()
// 	return nil
// }





