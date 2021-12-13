package source

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"

	"github.com/elchead/misou/search"
	"github.com/pkg/errors"
)
const AppConfigPath = "../appconfig.json" // TODO move

type BookmarkSearcher struct {
	bashPath string
}

func NewBookmarkSearcher(bookmarkBashPath string) *BookmarkSearcher {
	return &BookmarkSearcher{bookmarkBashPath}
}

func (b BookmarkSearcher) Search(query string) ([]search.SearchResult, error) {
	cmd := exec.Command(b.bashPath, query)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return []search.SearchResult{}, errors.Wrapf(err, "failed to run bookmark search from %s", b.bashPath)
	}
	reader := bufio.NewScanner(&out)
	res := []search.SearchResult{}
	for reader.Scan() {
		match := reader.Text()
		split := strings.Split(match, "\t")
		title := split[0]
		link := split[1]
		res = append(res, search.SearchResult{Title: title, Link: link, Content: "", Provider: "bookmark", ContentType: "txt"})
	}
	return res, nil
}
