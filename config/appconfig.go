package config

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/elchead/misou/filesystem"
	"github.com/elchead/misou/search"
	"github.com/elchead/misou/source"
	"github.com/pkg/errors"
)

type AppConfig struct {
	BookmarkBashSubPath string `json:"bookmarkBashSubPath"`
	FileScrapeRgaPath string `json:"fileScrapeRgaPath"`
	FileScrapePath      string `json:"fileScrapePath"`
	HistoryBashSubPath  string `json:"historyBashSubPath"`
	ReadwiseCsvPath     string `json:"readwiseCsvPath"`
	PeoplePath string `json:"peoplePath"`
	PeopleUrl string `json:"peopleUrl"`
	TwitterPath string `json:"twitterPath"`
	InstapaperPath     string `json:"instapaperPath"`
	RepoPath            string `json:"repoPath"`
	SecretsSubPath      string `json:"secretsSubPath"`
	Sources             struct {
		Bookmarks  bool `json:"bookmarks"`
		Gdrive     bool `json:"gdrive"`
		History    bool `json:"history"`
		Localfiles bool `json:"localfiles"`
		Readwise   bool `json:"readwise"`
		Instapaper bool `json:"instapaper"`
		People     bool `json:"people"`
		Twitter     bool `json:"twitter"`
	} `json:"sources"`
}

func (config AppConfig) InitSources(s search.SearcherI) {
	if config.Sources.Localfiles {
		fileSearcher := source.NewFileSearcher(config.FileScrapeRgaPath,config.FileScrapePath)
		s.AddApiClient(fileSearcher)
	}
	if config.Sources.Readwise {
		f, _ := filesystem.OpenFile(config.ReadwiseCsvPath)
		defer f.Close()
		readwiseClient := source.NewReadwiseCSVParser(f)
		s.AddSrcToDb(readwiseClient)
	}

	if config.Sources.People {
		f, err := filesystem.OpenFile(config.PeoplePath)
		if err != nil {
			log.Infof("Failed to open people file: %v", err)
		}
		peopleParser := source.NewPeopleParser(f)
		peopleParser.Url = config.PeopleUrl
		s.AddSrcToDb(peopleParser)
	}
	if config.Sources.Twitter {
		f, err := filesystem.OpenFile(config.TwitterPath)
		if err != nil {
			log.Infof("Failed to open people file: %v", err)
		}
		peopleParser := &source.TwitterParser{File:f}
		s.AddSrcToDb(peopleParser)
	}


	if config.Sources.Instapaper {
		f, _ := filesystem.OpenFile(config.InstapaperPath)
		defer f.Close()
		instapaperClient := source.NewInstapaperParser(f)
		s.AddSrcToDb(instapaperClient)
	}
	
	if config.Sources.Bookmarks {
		bookmarkBashPath := path.Join(config.RepoPath, config.BookmarkBashSubPath)
		bookmarkSearcher := source.NewBookmarkSearcher(bookmarkBashPath)
		s.AddApiClient(bookmarkSearcher)
	}

	if config.Sources.History {
		historyBashPath := path.Join(config.RepoPath, config.HistoryBashSubPath)
		historySearcher := source.NewBrowserHistorySearcher(historyBashPath)
		s.AddApiClient(historySearcher)
	}


	if config.Sources.Gdrive {
		secretPath := path.Join(config.RepoPath, config.SecretsSubPath)
		gdriveCredentialsPath := filepath.Join(secretPath, "gdrive.json")
		tokenPath := filepath.Join(secretPath, "token.json")

		credFile, err := os.Open(gdriveCredentialsPath)
		if err != nil {
			log.Fatal(err)
		}
		defer credFile.Close()

		fmt.Printf("Saving credential file to: %s\n", tokenPath)
		saveTokenFile, err := os.OpenFile(tokenPath, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalf("Unable to cache oauth token: %v", err)
		}
		defer saveTokenFile.Close()
		client := source.NewGdriveClient(secretPath, credFile, saveTokenFile, saveTokenFile)
		s.AddApiClient(client)
	}
}

func LoadConfig(filePath string) AppConfig {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	config, err := LoadConfigFromReader(file)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func LoadConfigFromReader(file io.Reader) (AppConfig, error) {
	var config AppConfig
	err := filesystem.LoadVariable(file, &config)
	if err != nil {
		return AppConfig{}, errors.Wrapf(err, "failed to load app config from file")
	}
	return config, nil
}

func (config AppConfig) BookmarkPath() string {
	return path.Join(config.RepoPath, config.BookmarkBashSubPath)
}

func (config AppConfig) HistoryPath() string {
	return config.RepoPath + config.HistoryBashSubPath
}
