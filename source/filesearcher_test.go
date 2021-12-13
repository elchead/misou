package source_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/elchead/misou/search"
	"github.com/elchead/misou/source"
	"github.com/stretchr/testify/assert"
)

func TestFileSearcher(t *testing.T) {
	folder := "/Users/adria/Google Drive/Obsidian/Second_brain"
	query := "antifragile"
	searcher := source.NewFileSearcher("/Users/adria/homebrew/bin/rga",folder)
	res, err := searcher.Search(query)
	assert.NoError(t, err)
	assert.NotEmpty(t, res[0].Title)
	assert.NotEmpty(t, res[0].Content)
	assert.Equal(t, true, len(res) > 1)
}

func TestResultParser(t *testing.T) {
	txt := `Books/Antifragile.md:title: "Antifragile: Things That Gain from Disorder - Nassim Nicholas Taleb"
Books/Antifragile.md:🔗 Link : [Goodreads](https://www.goodreads.com/book/show/13530973-antifragile)
Readwise/RW_Articles/Antifragile - By Nassim Nicholas Taleb  Derek Sivers.md:- Antifragile risk taking - not education and formal, organized research - is largely responsible for innovation and growth. ([View Highlight](https://instapaper.com/read/1443421199/17652560))
Journal/2021-10-08.md:-   Make today great:: todo`
res := source.ExtractContentAndMetadata(strings.NewReader(txt))
assert.Equal(t, "Books/Antifragile.md",res[0].Title)
assert.Equal(t, "Journal/2021-10-08.md",res[3].Title)
assert.Equal(t, "-   Make today great:: todo",res[3].Content)
}

func TestBatchDuplicateResult(t *testing.T) {
	res := []search.SearchResult{{Title: "A", Content: "A"}, {Title: "A", Content: "B"}}
	newRes := source.BatchDuplicateResults(res)
	assert.Equal(t, true, len(newRes) == 1)
	assert.NotEqual(t, "A", newRes[0].Content)
}

func TestAddScores(t *testing.T) {
	res := []search.SearchResult{{Title: "1", Content: "Finance lasfjds ksfdjsdfj jsfkddfjs jwkefw wjeew money"}, {Title: "2", Content: "Money is money and money will be money and money will be money and money will be money and money will be money"}}
	newRes := source.AddScores(res,"money")
	assert.Equal(t, true, newRes[0].Score > newRes[1].Score)
	fmt.Printf("%+v",newRes)
	assert.Equal(t, "2",newRes[0].Title )
}

func TestOpenWithObsidianIfMarkdown(t *testing.T) {
	file := ("Readwise/Hi hola.md")
	assert.Equal(t, "obsidian://open?file=Readwise%2FHi%20hola.md", source.ConstructLink(file))
}

func TestReadwiseMetaRemoved(t *testing.T) {
	t.Run("remove metadata", func(t *testing.T){
		txt := `# Brain Food: Some Useful Things - stobbe.adrian@gmail.com - Gmail
![rw-book-cover](https://readwise-assets.s3.amazonaws.com/static/images/article0.00998d930354.png)

## Metadata
- Author: [[mail.google.com]]
- Full Title: Brain Food: Some Useful Things - stobbe.adrian@gmail.com - Gmail
- Category: #articles
- URL: https://mail.google.com/mail/u/0/#inbox/FMfcgxwLtZrBpnTFVMGjWTjdbgMtTtlc

## Highlights
- You’re free when no one can buy your tim`
		got := source.GetReadwiseHighlights(txt)
		want := "- You’re free when no one can buy your tim" 
		assert.Equal(t, want, got)

		txt = `# The Psychology of Money

![rw-book-cover](https://readwise-assets.s3.amazonaws.com/static/images/article3.5c705a01b476.png)

## Metadata
- Author: [[notion.so]]
- Full Title: The Psychology of Money
- Category: #articles
- URL: https://www.notion.so/adrian99/The-Psychology-of-Money-9bcf55d7d14148108798be6490db9a27

## Highlights
- I want to spend money on what lets me spend my time in the most meaningful way. I am willing to pay for stuff I can do myself but is hard for me or dull, e.g. repairs, repetitive tasks
- Keep your expectations low. Wealth is not just about income but also about our spending behavior.`

	got = source.GetReadwiseHighlights(txt)
	want = `- I want to spend money on what lets me spend my time in the most meaningful way. I am willing to pay for stuff I can do myself but is hard for me or dull, e.g. repairs, repetitive tasks
- Keep your expectations low. Wealth is not just about income but also about our spending behavior.`
	assert.Equal(t, want, got)
	})

	t.Run("detect readwise result",func(t *testing.T) {
		title := "Readwise/RW_Articles/Antifragile - By Nassim Nicholas Taleb Derek Sivers.md"
		assert.Equal(t, true, source.IsReadwiseResult(title))
	})
}
