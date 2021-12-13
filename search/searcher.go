package search

import (
	"sort"
	"sync"

	log "github.com/sirupsen/logrus"
)

type SearcherI interface {
	AddApiClient(DataApi)
	AddSrcToDb(DbSource)
	Search(string) ([]SearchResult, error)
}


type Indexer interface {
	AddSrcToDb([]*IndexData)
	Search(string) ([]SearchResult, error)
}

type Searcher struct {
	apiClients  []DataApi
	indexer Indexer
}

func NewSearcher() *Searcher {
	return &Searcher{apiClients: []DataApi{},indexer: NewBleveIndexer()}
}

func NewSearcherUsingIndexer(idx Indexer) *Searcher {
	return &Searcher{apiClients: []DataApi{},indexer: idx}
}

func (s *Searcher) AddApiClient(src DataApi) {
	s.apiClients = append(s.apiClients, src)
}


func (s *Searcher) AddSrcToDb(src DbSource) {
	data, err := src.TransformToData()
	if err != nil {
		log.Infof("Failed to transform data: %v. Skipping source",err)
		return
	}
	s.indexer.AddSrcToDb(data)
}

func (s Searcher) searchApis(keywords string) ([]SearchResult, error) {
	var totalRes []SearchResult
	if !containsMeaningfulQuery(keywords) {
		return totalRes, nil
	}
	return runParallelApiQueries(keywords,s.apiClients), nil
}

func runParallelApiQueries(keywords string,apiClients []DataApi) ([]SearchResult) {
	var wg sync.WaitGroup
	var resChan chan []SearchResult = make(chan []SearchResult, len(apiClients))
	for _, client := range apiClients {
		wg.Add(1)
		go func(wg *sync.WaitGroup,res chan []SearchResult,client DataApi) {
			defer wg.Done()
			gres, err := client.Search(keywords)
			if err == nil {
				res <- gres
			}	
		}(&wg,resChan,client)
	}
	wg.Wait()
	close(resChan)	
	var totalRes []SearchResult
	for val := range resChan {
		totalRes = append(totalRes, val...)
	}
	return totalRes
}

func (s Searcher) Search(keywords string) ([]SearchResult, error) {
	var totalRes []SearchResult
	apiRes, err := s.searchApis(keywords)
	if err != nil {
		return []SearchResult{}, err
	}
	totalRes = append(totalRes, apiRes...)

	indexerRes, _ := s.indexer.Search(keywords)
	totalRes = append(totalRes, indexerRes...)	
	sort.Slice(totalRes, func(i, j int) bool {
		return totalRes[i].Score > totalRes[j].Score
	      })
	return totalRes, nil
}

func containsMeaningfulQuery(query string) bool {
	return len(Analyze(query)) > 0
}
