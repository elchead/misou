package config_test

import (
	"fmt"
	"os"
	"path"
	"reflect"
	"testing"

	"github.com/elchead/misou/config"
	"github.com/elchead/misou/search"
	"github.com/elchead/misou/source"
	"github.com/stretchr/testify/assert"
)

func findInSources(sources []interface{}) map[string]bool {
	found := map[string]bool{}
	for _, src := range sources {
		fmt.Println(reflect.TypeOf(src))
		switch src.(type) {
		case *source.BookmarkSearcher:
			found["bookmark"] = true
			// assert.Equal(t, true, config.Sources.Bookmarks)
		case *source.ReadwiseCSVParser:
			found["readwise"] = true
		case *source.BrowserHistorySearcher:
			found["history"] = true
		case *source.FileSearch:
			found["file"] = true
		case *source.GoogleDriveClient:
			found["gdrive"] = true
		case *source.InstapaperParser:
			found["instapaper"] = true
		case *source.PeopleParser:
			found["people"] = true // TODO use provider constants
		}
	}
	return found
}

func assertOnlySourcesForConfigAdded(t testing.TB, config *config.AppConfig, s *fakeSearcher) {
	t.Helper()
	foundMap := findInSources(s.Sources)
	assert.Equalf(t, config.Sources.Bookmarks, foundMap["bookmark"], "Bookmark wrongly configured")
	assert.Equalf(t, config.Sources.Readwise, foundMap["readwise"], "Readwise wrongly configured")
	assert.Equalf(t, config.Sources.History, foundMap["history"], "History wrongly configured")
	assert.Equalf(t, config.Sources.Localfiles, foundMap["file"], "File wrongly configured")
	assert.Equalf(t, config.Sources.Gdrive, foundMap["gdrive"], "Gdrive wrongly configured")
	assert.Equalf(t, config.Sources.Instapaper, foundMap["instapaper"], "Instapaper wrongly configured")
	assert.Equalf(t, config.Sources.People, foundMap["people"], "People wrongly configured")
}

func TestLoadConfig(t *testing.T) {
	repoPath := ".."
	file, err := os.Open(path.Join(repoPath, "appconfig.json"))
	assert.NoError(t, err)
	config, err := config.LoadConfigFromReader(file)
	assert.NoError(t, err)
	assert.NotEmpty(t, config)
	t.Run("read activated sources", func(t *testing.T) {
		assert.NotEmpty(t, config.Sources)
	})
}

type fakeSearcher struct {
	Sources []interface{}
}

func NewFakeSearcher() *fakeSearcher {
	return &fakeSearcher{Sources: make([]interface{}, 0)}
}

func (f *fakeSearcher) AddApiClient(c search.DataApi) {
	f.Sources = append(f.Sources, c)
}

func (f *fakeSearcher) AddSrcToDb(c search.DbSource) {
	f.Sources = append(f.Sources, c)
}

func (f *fakeSearcher) Search(string) ([]search.SearchResult, error) {
	return []search.SearchResult{}, nil
}
func TestInitSources(t *testing.T) {
	cfg := config.AppConfig{Sources: struct {
		Bookmarks  bool `json:"bookmarks"`
		Gdrive     bool `json:"gdrive"`
		History    bool `json:"history"`
		Localfiles bool `json:"localfiles"`
		Readwise   bool `json:"readwise"`
		Instapaper     bool `json:"instapaper"`
		People     bool `json:"people"`
		Twitter bool `json:"twitter"`
	}{Bookmarks: true, Gdrive: false, History: true, Localfiles: true, Readwise: true,Instapaper: true, People: true, Twitter: true},
		ReadwiseCsvPath: "/Users/adria/Programming/misou/data/readwise-data.csv",InstapaperPath: "/Users/adria/Programming/misou/data/instapaper-full-text.json",PeoplePath: "/Users/adria/Programming/misou/data/people.json",TwitterPath:"/Users/adria/Programming/misou/data/tweet.json",} // TODO remove file logic inside InitSources (use mock..)
	searcher := NewFakeSearcher()
	cfg.InitSources(searcher)
	assertOnlySourcesForConfigAdded(t, &cfg, searcher)
}

func TestBookmarkPath(t *testing.T) {
	cfg := config.AppConfig{RepoPath: "/Users/adria/Programming/misou",BookmarkBashSubPath: "/source/bookmarks.sh"}
	assert.Equal(t,"/Users/adria/Programming/misou/source/bookmarks.sh",cfg.BookmarkPath())
}
//TODO func TestCatchMissingParameterConfig(t *testing.T) {}
