package source

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"

	"github.com/elchead/misou/search"
	"github.com/pkg/errors"
)

type BrowserHistorySearcher struct {
	bashPath string
}

// TODO sometimes links have the same title but different pages (save different links in content?)
func removeDuplicateTitles(slice []search.SearchResult) []search.SearchResult {
	visited := make(map[string]bool)
	list := []search.SearchResult{}
	for _, entry := range slice {
		if _, value := visited[entry.Title]; !value {
			visited[entry.Title] = true
			list = append(list, entry)
		}
	}
	return list
}

func NewBrowserHistorySearcher(bashScriptPath string) *BrowserHistorySearcher {
	return &BrowserHistorySearcher{bashScriptPath}
}

func (b BrowserHistorySearcher) Search(query string) ([]search.SearchResult, error) {
	cmd := exec.Command(b.bashPath, query)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return []search.SearchResult{}, errors.Wrapf(err, "failed to run history search from %s", b.bashPath)
	}
	reader := bufio.NewScanner(&out)
	res := []search.SearchResult{}
	for reader.Scan() {
		match := reader.Text()
		split := strings.Split(match, ";")
		title := split[0]
		link := split[1]
		res = append(res, search.SearchResult{Title: title, Link: link, Content: "", Provider: "history", ContentType: "txt"})
	}
	res = removeDuplicateTitles(res)
	return res, nil
}
