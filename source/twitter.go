package source

import (
	"os"

	"github.com/elchead/misou/filesystem"
	"github.com/elchead/misou/search"
)

const twitterProvider = "twitter"
const twitterUrl = "https://twitter.com/astobbe_/status/"
type TwitterParser struct {
	File *os.File
}

type tweet struct {
	CreatedAt        string   `json:"created_at"`
	DisplayTextRange []string `json:"display_text_range"`
	Entities         struct {
		Hashtags []interface{} `json:"hashtags"`
		Media    []struct {
			DisplayURL    string   `json:"display_url"`
			ExpandedURL   string   `json:"expanded_url"`
			ID            string   `json:"id"`
			IdStr         string   `json:"id_str"`
			Indices       []string `json:"indices"`
			MediaURL      string   `json:"media_url"`
			MediaUrlHttps string   `json:"media_url_https"`
			Sizes         struct {
				Large struct {
					H      string `json:"h"`
					Resize string `json:"resize"`
					W      string `json:"w"`
				} `json:"large"`
				Medium struct {
					H      string `json:"h"`
					Resize string `json:"resize"`
					W      string `json:"w"`
				} `json:"medium"`
				Small struct {
					H      string `json:"h"`
					Resize string `json:"resize"`
					W      string `json:"w"`
				} `json:"small"`
				Thumb struct {
					H      string `json:"h"`
					Resize string `json:"resize"`
					W      string `json:"w"`
				} `json:"thumb"`
			} `json:"sizes"`
			SourceStatusID    string `json:"source_status_id"`
			SourceStatusIdStr string `json:"source_status_id_str"`
			SourceUserID      string `json:"source_user_id"`
			SourceUserIdStr   string `json:"source_user_id_str"`
			Type              string `json:"type"`
			URL               string `json:"url"`
		} `json:"media"`
		Symbols      []interface{} `json:"symbols"`
		Urls         []interface{} `json:"urls"`
		UserMentions []struct {
			ID         string   `json:"id"`
			IdStr      string   `json:"id_str"`
			Indices    []string `json:"indices"`
			Name       string   `json:"name"`
			ScreenName string   `json:"screen_name"`
		} `json:"user_mentions"`
	} `json:"entities"`
	ExtendedEntities struct {
		Media []struct {
			DisplayURL    string   `json:"display_url"`
			ExpandedURL   string   `json:"expanded_url"`
			ID            string   `json:"id"`
			IdStr         string   `json:"id_str"`
			Indices       []string `json:"indices"`
			MediaURL      string   `json:"media_url"`
			MediaUrlHttps string   `json:"media_url_https"`
			Sizes         struct {
				Large struct {
					H      string `json:"h"`
					Resize string `json:"resize"`
					W      string `json:"w"`
				} `json:"large"`
				Medium struct {
					H      string `json:"h"`
					Resize string `json:"resize"`
					W      string `json:"w"`
				} `json:"medium"`
				Small struct {
					H      string `json:"h"`
					Resize string `json:"resize"`
					W      string `json:"w"`
				} `json:"small"`
				Thumb struct {
					H      string `json:"h"`
					Resize string `json:"resize"`
					W      string `json:"w"`
				} `json:"thumb"`
			} `json:"sizes"`
			SourceStatusID    string `json:"source_status_id"`
			SourceStatusIdStr string `json:"source_status_id_str"`
			SourceUserID      string `json:"source_user_id"`
			SourceUserIdStr   string `json:"source_user_id_str"`
			Type              string `json:"type"`
			URL               string `json:"url"`
		} `json:"media"`
	} `json:"extended_entities"`
	FavoriteCount     string `json:"favorite_count"`
	Favorited         bool   `json:"favorited"`
	FullText          string `json:"full_text"`
	ID                string `json:"id"`
	IdStr             string `json:"id_str"`
	Lang              string `json:"lang"`
	PossiblySensitive bool   `json:"possibly_sensitive"`
	RetweetCount      string `json:"retweet_count"`
	Retweeted         bool   `json:"retweeted"`
	Source            string `json:"source"`
	Truncated         bool   `json:"truncated"`
}

type twitterEntry struct {
	Tweet tweet `json:"tweet"`
}

func (t TwitterParser) TransformToData() ([]*search.IndexData,error) {
	var entries []twitterEntry
	err := filesystem.LoadVariable(t.File,&entries)
	if err != nil {
		return []*search.IndexData{},err
	}
	// fmt.Println(entries)
	// return []*search.IndexData{},nil
	var indexData []*search.IndexData
	for _, entry := range entries {
		data := entry.Tweet
		// fmt.Println(t.transformTweet(data))
		indexData = append(indexData,t.transformTweet(data))
	}
	
	return indexData,nil
}

func (t TwitterParser) transformTweet(data tweet) *search.IndexData {
	var url string
	if len(data.Entities.Media) > 0 {
		url = data.Entities.Media[0].URL
	} else {
		url = twitterUrl + data.ID
			
	}
	content := data.FullText
	return &search.IndexData{
		Title: "",
		Link: url,
		Content: content,
		Provider: twitterProvider,
	}
}

