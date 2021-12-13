package source

import (
	"fmt"
	"os"
	"reflect"

	"github.com/elchead/misou/filesystem"
	"github.com/elchead/misou/search"
)

type PeopleParser struct {
	File *os.File
	Url string
}

func NewPeopleParser(f *os.File) *PeopleParser {
	return &PeopleParser{File:f}
}

type peopleEntry struct {
	Ask                                    string `json:"Ask"`
	Family                                 string `json:"Family"`
	Groups                                 string `json:"Groups"`
	How_Can_I_Help_                        string `json:"How Can I Help?"`
	Interests                              string `json:"Interests"`
	Last_Update                            string `json:"Last Update"`
	Likes                                  string `json:"Likes"`
	Name                                   string `json:"Name"`
	Website                                string `json:"Website"`
	Work                                   string `json:"Work"`
}

const peopleProvider  = "people"

func prettyPrintPerson(p peopleEntry) string {
	res := ""
	v := reflect.ValueOf(p)
	typeOfS := v.Type()
	for i := 0; i< v.NumField(); i++ {
		value := v.Field(i)
		if value.IsValid() && value.Interface() != "" {
			res += fmt.Sprintf("%s: %s\n",typeOfS.Field(i).Name, value.Interface())
		}
	}
	// res += fmt.Sprintf("Name: %s\n", p.Name)
	// res += fmt.Sprintf("Family: %s\n", p.Family)
	// res += fmt.Sprintf("Groups: %s\n", p.Groups)
	// res += fmt.Sprintf("Interests: %s\n", p.Interests)
	// res += fmt.Sprintf("Likes: %s\n", p.Likes)
	// res += fmt.Sprintf("Work: %s\n", p.Work)
	// res += fmt.Sprintf("Website: %s\n", p.Website)
	// res += fmt.Sprintf("How Can I Help?: %s\n", p.How_Can_I_Help_)
	// res += fmt.Sprintf("Last Update: %s\n", p.Last_Update)
	return res
}

func (p PeopleParser) constructUrl(name string) string {
	return fmt.Sprintf("%s/?q=%s", p.Url,name)
}

func (p PeopleParser) transformPeopleEntries(entries []peopleEntry) []*search.IndexData {
	res := make([]*search.IndexData, len(entries))
	for i,entry := range entries {
		content := prettyPrintPerson(entry)
		res[i] = &search.IndexData{Title: entry.Name,Link: p.constructUrl(entry.Name), Content: content, Provider: peopleProvider, ContentType: "contact"}
	}
	return res
}



func (p PeopleParser) TransformToData() ([]*search.IndexData, error) {
	var entries []peopleEntry
	err := filesystem.LoadVariable(p.File,&entries)
	if err != nil {
		return []*search.IndexData{},err
	}
	return p.transformPeopleEntries(entries),nil
}
