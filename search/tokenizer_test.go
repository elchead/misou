package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(t *testing.T) {
	data := "Hi world, I am Adrian and he wants to learn Golang."
	tokens := Analyze(data)
	assert.Equal(t,[]string{"hi","world","adrian","want","learn","golang"},tokens)
}

func TestDetectTrivialWords(t *testing.T) {
	assert.False(t,containsMeaningfulQuery("to"))
	assert.False(t,containsMeaningfulQuery("I"))
}


