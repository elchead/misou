package search

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	fs "github.com/elchead/misou/filesystem"
	jsoniter "github.com/json-iterator/go"
)

//maps tokens to array of record ids
var indexer tokenToIdx
var records mapIdToRecord

func ResetMemoryDb() {
	indexer = make(tokenToIdx)
	records = make(mapIdToRecord)	
}

func init() {
	ResetMemoryDb()
}

type mapIdToRecord map[Id]Record
type tokenToIdx map[string][]Id

// refers to record id
func getId() Id {
	len := len(records)
	return Id(strconv.Itoa(len))
}

func AddEntry(src DbSource){
	entries,_ := src.TransformToData()
	for _, e := range entries {
		record := dataToRecord(e,getId())
		registerRecord(record)
	}
}

func Add(entry *IndexData){
	record := dataToRecord(entry,getId())
	registerRecord(record)	
}


// OR match
func Search(query string) ([]SearchResult,error) {
	queries := Analyze(query)
	results := make(map[Id]bool)
	if len(queries) == 0 {
		return []SearchResult{}, errors.New("empty query")
	}
	for _, query := range queries {
		recordsWithQuery := indexer[query]
		for _, recordID := range recordsWithQuery {
			_, inMap := results[recordID]
			if !inMap {
				results[recordID] = true
			}
		}
	}
	currentSearchResults := make(map[string]string)
	return rank(results, queries, currentSearchResults),nil
}

func registerRecord(record Record) {
	records[record.ID] = record
}

func GenerateIndexer() {
	flushSavedRecordsIntoInvertedIndex(records)	
}

func dataToRecord(d *IndexData,uid Id) Record {
	//tokenize, stem, and filter
	tokens := Analyze(d.Content)

	//count frequency and create `Record`
	frequencyOfTokens := countFrequencyTokens(tokens)

	//adds meta level tags defined into the data - how do we set the frequency? Since these are global tags
	//we push some more probability on them since the user said these were important to index by
	//use a simple heuristic of pushing ~20% of "counts" on them
	//TODO: is there a more intellignet heuristic we can use here
	frequencyToAdd := len(tokens) / 5
	for _, metaTag := range d.Tags {
		_, metaTagInMap := frequencyOfTokens[metaTag]
		if metaTagInMap {
			frequencyOfTokens[metaTag] += frequencyToAdd
		} else {
			frequencyOfTokens[metaTag] = frequencyToAdd
		}
	}

	//store record in our tokens list
	var provider string
	if d.Provider == "" {
		provider = fileProvider
	} else { provider = d.Provider }
	record := Record{ID: uid, Title: d.Title, Link: d.Link, Content: d.Content, TokenFrequency: frequencyOfTokens, Provider: provider}
	return record
}

func LoadDb(recordsFile io.Reader,indexFile io.Reader) error {
	err := fs.LoadVariable(recordsFile,&records)
	if err != nil {
		return fmt.Errorf("could not load db: %v", err)
	}
	return fs.LoadVariable(indexFile,&indexer)
}

func WriteRecordsToFile(input io.Writer) {
	jsoniter.NewEncoder(input).Encode(records)
}

func SaveDb(recFile io.Writer,idxFile io.Writer) error {
	err := fs.WriteVariable(recFile,records)
	if err != nil {
		return fmt.Errorf("could not save db: %v", err)
	}
	return fs.WriteVariable(idxFile,indexer)
}

func readRecords(file io.Reader) mapIdToRecord {
	var got mapIdToRecord
	jsoniter.NewDecoder(file).Decode(&got) // TODO error handling
	return got
}

func ReadRecord(id Id) (Record,error) {
	val, ok := records[id]
	if !ok {
		return Record{}, fmt.Errorf("Record was not found")
	}
	return val, nil
}

func ReadRecordFromFile(file io.Reader,id Id) (Record, error) {
	records := readRecords(file)
	val,ok := records[id]
	if !ok {
		return Record{}, fmt.Errorf("Record was not found")
	}
	return val, nil
}
