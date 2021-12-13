package search

import (
	"fmt"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/blevesearch/bleve/v2"
	index "github.com/blevesearch/bleve_index_api"
)


type OpenedBleveIndexer struct {
	indexer bleve.Index
	uid int
}

type BleveIndexer struct {
	indexer bleve.Index
	uid int
}

func (s *OpenedBleveIndexer) AddSrcToDb(entries []*IndexData)  {
}

func (s *OpenedBleveIndexer) ExactSearch(keywords string) ([]SearchResult, error) {
	query := bleve.NewMatchPhraseQuery(keywords)
	search := bleve.NewSearchRequest(query)
	// search.Highlight = bleve.NewHighlight() // without highlight no text match is returned...
	searchResults, err := s.indexer.Search(search)

	var results []SearchResult
	for _,hit := range searchResults.Hits {
		id := hit.ID
		doc, _ := s.indexer.Document(id)
		var res SearchResult
		doc.VisitFields(func(fo index.Field) { 
			// fmt.Printf("%s:%s\n",fo.Name(),fo.Value())
			switch fo.Name() {
			case "content":
				res.Content = string(fo.Value())
			case "title":
				res.Title = string(fo.Value())
			
			case "link":
				res.Link = string(fo.Value())
			case "contentType":
				res.ContentType = string(fo.Value())
			case "provider":
				res.Provider = string(fo.Value())
			}
		})
		results = append(results,res)

	}
	return results,err
}
func (s *OpenedBleveIndexer) Search(keywords string) ([]SearchResult, error) {
	// defer os.RemoveAll(bleveDir)
	if isExactQuery(keywords) {
		return s.ExactSearch(keywords)
	}
	query := bleve.NewMatchQuery(keywords)
	search := bleve.NewSearchRequest(query)
	// search.Highlight = bleve.NewHighlight() // without highlight no text match is returned...
	searchResults, err := s.indexer.Search(search)

	var results []SearchResult
	for _,hit := range searchResults.Hits {
		id := hit.ID
		doc, _ := s.indexer.Document(id)
		var res SearchResult
		doc.VisitFields(func(fo index.Field) { 
			// fmt.Printf("%s:%s\n",fo.Name(),fo.Value())
			switch fo.Name() {
			case "content":
				res.Content = string(fo.Value())
			case "title":
				res.Title = string(fo.Value())
			
			case "link":
				res.Link = string(fo.Value())
			case "contentType":
				res.ContentType = string(fo.Value())
			case "provider":
				res.Provider = string(fo.Value())
			}
		})
		res.Score = hit.Score
		results = append(results,res)

	}
	return BatchDuplicateResults(results),err
}

func NewPersistentBleveIndexer(bleveDir string) Indexer {
	mapping := bleve.NewIndexMapping()
	// blogM := bleve.NewDocumentMapping()
	// mapping.DefaultMapping
	// mapping.AddDocumentMapping("blog",blogM)
	fmt.Println("trying to create index")
	index, err := bleve.New(bleveDir, mapping) // TODO
	if err != nil {
		fmt.Println("folder exists, trying to open index")
		index, err = bleve.Open(bleveDir)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("opened existing index")
		return &OpenedBleveIndexer{indexer: index,uid:1}
	}
	return &BleveIndexer{indexer: index,uid:1}
}

func OpenBleveIndexer(bleveDir string) *BleveIndexer {
	index, err := bleve.Open(bleveDir)
	if err != nil {
		log.Fatal(err)
	}
	return &BleveIndexer{indexer: index, uid:1}
}

func NewBleveIndexer() *BleveIndexer {
	mapping := bleve.NewIndexMapping()
	// blogM := bleve.NewDocumentMapping()
	// mapping.DefaultMapping
	// mapping.AddDocumentMapping("blog",blogM)
	index, err := bleve.NewUsing("", mapping, bleve.Config.DefaultIndexType,
	bleve.Config.DefaultMemKVStore, nil) //New(bleveDir, mapping) // TODO
	if err != nil { log.Fatal(err) }
	return &BleveIndexer{indexer: index,uid:1}
}

func (s *BleveIndexer) AddSrcToDb(entries []*IndexData) {
	batchSz := 97
	if len(entries)< batchSz { batchSz = 1 }
	batch := s.indexer.NewBatch()
	for idx, entry := range entries {
		id := strconv.Itoa(s.uid)
		if err:=batch.Index(id,*entry);err!=nil { log.Infof("Entry %v could not be added: %v",*entry,err) }
		s.uid += 1
		isNewBatch := (idx % batchSz) == 0
		if isNewBatch {
			s.indexer.Batch(batch)
			batch = s.indexer.NewBatch()
		}
	}
	for _,entry := range entries[len(entries)-batchSz:] {
		id := strconv.Itoa(s.uid)
		if err:=batch.Index(id,*entry);err!=nil { log.Infof("Entry %v could not be added: %v",*entry,err) }
		s.uid += 1
	}
	s.indexer.Batch(batch)
}



func (s *BleveIndexer) AddEntry(entry *IndexData) {
	id := strconv.Itoa(s.uid)
	err := s.indexer.Index(id,*entry)
	if err != nil {
		log.Infof("Entry %v could not be added: %v",*entry,err)
	}
	s.uid += 1
}

func (s *BleveIndexer) ExactSearch(keywords string) ([]SearchResult, error) {
	query := bleve.NewMatchPhraseQuery(keywords)
	search := bleve.NewSearchRequest(query)
	// search.Highlight = bleve.NewHighlight() // without highlight no text match is returned...
	searchResults, err := s.indexer.Search(search)

	var results []SearchResult
	for _,hit := range searchResults.Hits {
		id := hit.ID
		doc, _ := s.indexer.Document(id)
		var res SearchResult
		doc.VisitFields(func(fo index.Field) { 
			// fmt.Printf("%s:%s\n",fo.Name(),fo.Value())
			switch fo.Name() {
			case "content":
				res.Content = string(fo.Value())
			case "title":
				res.Title = string(fo.Value())
			
			case "link":
				res.Link = string(fo.Value())
			case "contentType":
				res.ContentType = string(fo.Value())
			case "provider":
				res.Provider = string(fo.Value())
			}
		})
		results = append(results,res)

	}
	return results,err
}
func (s *BleveIndexer) Search(keywords string) ([]SearchResult, error) {
	// defer os.RemoveAll(bleveDir)
	if isExactQuery(keywords) {
		return s.ExactSearch(keywords)
	}
	query := bleve.NewMatchQuery(keywords)
	search := bleve.NewSearchRequest(query)
	// search.Highlight = bleve.NewHighlight() // without highlight no text match is returned...
	searchResults, err := s.indexer.Search(search)

	var results []SearchResult
	for _,hit := range searchResults.Hits {
		id := hit.ID
		doc, _ := s.indexer.Document(id)
		var res SearchResult
		doc.VisitFields(func(fo index.Field) { 
			// fmt.Printf("%s:%s\n",fo.Name(),fo.Value())
			switch fo.Name() {
			case "content":
				res.Content = string(fo.Value())
			case "title":
				res.Title = string(fo.Value())
			
			case "link":
				res.Link = string(fo.Value())
			case "contentType":
				res.ContentType = string(fo.Value())
			case "provider":
				res.Provider = string(fo.Value())
			}
		})
		res.Score = hit.Score
		results = append(results,res)

	}
	return BatchDuplicateResults(results),err
}

func batchSearchResult(newResult SearchResult, results *[]SearchResult) []SearchResult {
	for i := range *results {
		if (*results)[i].Title == newResult.Title{
			(*results)[i].Content += "\n--------------------------------\n"
			(*results)[i].Content += newResult.Content
		}
	}
	return *results
}

type visitorCheck map[string]map[string]bool 

func (v *visitorCheck) isVisited(title,provider string) bool {
	if((*v)[title] == nil){
		(*v)[title] = make(map[string]bool)
	}
	_, visitedBefore := (*v)[title][provider]
	return visitedBefore	
}

func (v *visitorCheck) Visit(title,provider string) {
	(*v)[title][provider] = true	
}

func BatchDuplicateResults(slice []SearchResult) []SearchResult {
	visitor := make(visitorCheck)
	list := []SearchResult{}
	for _, entry := range slice {
		if visitedBefore := visitor.isVisited(entry.Title,entry.Provider); !visitedBefore {
			
			visitor.Visit(entry.Title,entry.Provider)
			list = append(list, entry)
		} else {
			batchSearchResult(entry, &list)
		}
	}
	return list
}

func isExactQuery(query string) bool {
	return strings.Contains(query,"\"")
}

func removeQuotes(s string) string {
	return strings.Trim(s, "\"")
}
