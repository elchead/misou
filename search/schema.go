package search

type IndexData struct {
	Title   string   `json:"title"`
	Link    string   `json:"link"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	Provider    string `json:"provider"`
	ContentType string `json:"contentType"`
}

type Id string

// database unit for all sources
type Record struct {
	//unique identifier
	ID Id `json:"id"`
	//title
	Title string `json:"title"`
	//potential link to the source if applicable
	Link string `json:"link"`
	//text content to display on results page
	Content string `json:"content"`
	//map of tokens to their frequency
	TokenFrequency map[string]int `json:"tokenFrequency"`
	Provider string `json:"provider"`
}

// type SearchResult struct {
// 	IndexData
// 	Matches     int    `json:"matches"`
// }
type SearchResult struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Content     string `json:"content"`
	Provider    string `json:"provider"`
	Matches     int    `json:"matches"`
	ContentType string `json:"contentType"`
	Score       float64 `json:"score"`
}

type DummyFile struct {
	Content string
	Title   string
}

func (d DummyFile) TransformToData() ([]*IndexData,error) {
	return []*IndexData{{Title: d.Title, Content: d.Content}},nil
}
