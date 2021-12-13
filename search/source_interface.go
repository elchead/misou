package search

type DataApi interface {
	Search(string) ([]SearchResult, error)
}

type DbSource interface {
	TransformToData() ([]*IndexData, error)
}
