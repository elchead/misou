package search

//takes a string of tokens and returns a map of each token to its frequency
func countFrequencyTokens(tokens []string) map[string]int {
	frequencyWords := make(map[string]int)
	for _, token := range tokens {
		_, isInMap := frequencyWords[token]
		if isInMap {
			frequencyWords[token] += 1
		} else {
			frequencyWords[token] = 1
		}
	}
	return frequencyWords
}

//takes all of the saved records and puts them in our inverted index
func flushSavedRecordsIntoInvertedIndex(recordList map[Id]Record) {
	//we already have token frequency data precomputed and saved, so just add it to inverted index directly
	for key, record := range recordList {
		writeTokenFrequenciesToInvertedIndex(record.TokenFrequency, key)
	}
}

//write a map of tokens to their counts in our inverted index
func writeTokenFrequenciesToInvertedIndex(frequencyOfTokens map[string]int, uniqueID Id) {
	//loop through final frequencyOfTokens and add it to our inverted index database
	for key, _ := range frequencyOfTokens {
		_, keyInInvertedIndex := indexer[key]
		if keyInInvertedIndex {
			indexer[key] = append(indexer[key], uniqueID)
		} else {
			indexer[key] = []Id{uniqueID}
		}
	}
}
