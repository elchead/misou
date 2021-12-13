package search

type ApolloIndexer struct {
	initialized bool
}

func NewApolloIndexer() *ApolloIndexer {
	return &ApolloIndexer{initialized: false}
}

func (s *ApolloIndexer) AddSrcToDb(data []*IndexData) {
	s.AddEntries(data)                 // Move Db to search pkg
}

func (s *ApolloIndexer) AddEntries(entries []*IndexData) {
	for _, entry := range entries {
		Add(entry)
	}
}


func (s *ApolloIndexer) AddEntry(entry *IndexData) {
	record := dataToRecord(entry,getId())
	registerRecord(record)	
}

func (s *ApolloIndexer) init() {
	s.initialized = true
	GenerateIndexer()
}


func (s *ApolloIndexer) Search(keywords string) ([]SearchResult, error) {
	if !s.initialized {
		s.init()
		s.initialized = true	
	}
	return Search(keywords) // TODO remove function from pkg
}

