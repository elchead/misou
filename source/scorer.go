package source

import (
	"math"
	"regexp"
	"strings"
)

func termFrequency(term,txt string) int {
	term = strings.ToLower(term)
	txt = strings.ToLower(txt)
	re := regexp.MustCompile(term)
	matches := re.FindAllStringIndex(txt, -1)
	return len(matches)
}

func fieldNormalization(txt string) float64 {
	return 1/math.Sqrt(float64(len(strings.Split(txt," "))))
}

func inverseDocumentFrequency(docCount, nbrResults int) float64 {
	return 1 + math.Log(float64(docCount)/float64(nbrResults + 1))
}

func getGlobalScore(term, txt string) float64 {
	tf := termFrequency(term,txt)
	norm := fieldNormalization(txt)
	return float64(tf) * norm //  * idf 
}
