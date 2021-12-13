package source

import (
	"bufio"
	"bytes"
	"io"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	"github.com/elchead/misou/search"
	log "github.com/sirupsen/logrus"
)

type FileSearch struct {
	Rga string
	Folder string
}

const fileSearchProvider = "file"

func batchSearchResult(newResult search.SearchResult, results *[]search.SearchResult) []search.SearchResult {
	for i := range *results {
		if (*results)[i].Title == newResult.Title {
			(*results)[i].Content += "\n--------------------------------\n"
			(*results)[i].Content += newResult.Content
		}
	}
	return *results
}

func BatchDuplicateResults(slice []search.SearchResult) []search.SearchResult {
	visited := make(map[string]bool)
	list := []search.SearchResult{}
	for _, entry := range slice {
		if _, visitedBefore := visited[entry.Title]; !visitedBefore {
			visited[entry.Title] = true
			list = append(list, entry)
		} else {
			batchSearchResult(entry, &list)
		}
	}
	return list
}

func formatObsidianPath(path string) string {
	replacedSpace := strings.Replace(path," ","%20",-1)
	return strings.Replace(replacedSpace,"/","%2F",-1)
}

func ConstructLink(path string) string {
	link := ""
	if filepath.Ext(path) == ".md" {
		link = "obsidian://open?file=" + formatObsidianPath(path)
	}
	// "file:///" + f.Folder + "/" + file
	return link//url.QueryEscape(link)
}

func NewFileSearcher(rgaPath, pathFolder string) *FileSearch {
	return &FileSearch{Rga: rgaPath, Folder: pathFolder}
}

func (f FileSearch) Search(query string) ([]search.SearchResult, error) {
	cmd := exec.Command(f.Rga, "-S","-w",query) 
	// -S (Searches case insensitively if the pattern is all lowercase. Search case sensitively otherwise.)
	//DISABLE? -o, --only-matching
	// Print only the matched (non-empty) parts of a matching line, with each such
	// part on a separate output line.
	// consider json: (highlighting?){"type":"match","data":{"path":{"text":"Blog/AMM-DeFi.md"},"lines":{"text":"Got excited about AMM in how it allows everyone to participate in markets and services that had been reserved to big liquidity providers. Users no longer trade against counterparties, but instead trade against liquidity locked inside smart smart contracts. Instead of believing in a central intermediary, we just need to trust a smart contract, which is transparent and (not manipulable?). The contracts form liquidity pools, which assure that the commodities e.g. curriencies can be exchanged. The exciting part for the John Doe like us is that we can get rendites from providing our liquidity in this pool. As such we are rewarded with a fraction of the transaction fees and by depositing an appropriate ratio of digital asses in the pool, you can further benefit from yield farming rewards. The idea behind it is to stabilize the market to avoid slippage (price difference when no buyer can be found quickly). This is better than P2P lending and not too mention lending to a bank. You participate directly, without intermediaries!"},"line_number":11,"absolute_offset":201,"submatches":[{"match":{"text":"AMM"},"start":18,"end":21},{"match":{"text":"liquidity"},"start":117,"end":126},{"match":{"text":"liquidity"},"start":210,"end":219},{"match":{"text":"liquidity"},"start":409,"end":418},{"match":{"text":"liquidity"},"start":585,"end":594}]}}

	cmd.Dir = f.Folder

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	res := ExtractContentAndMetadata(&out)
	res = BatchDuplicateResults(res)
	res = AddScores(res, query)
	return res, nil
}

func ExtractContentAndMetadata(out io.Reader) []search.SearchResult {
	reader := bufio.NewScanner(out)
	res := []search.SearchResult{}
	for reader.Scan() {
		match := reader.Text()
		split := strings.Split(match, ":")
		filename := split[0]
		txt := strings.TrimPrefix(match,filename+":")
		// if IsReadwiseResult(filename) { // rga does not read complete file content!
		// 	txt = GetReadwiseHighlights(txt)
		// }
		res = append(res, search.SearchResult{Title: filename, Link: ConstructLink(filename), Content: txt, Provider: fileSearchProvider, ContentType: "txt"})

	}
	return res
}

func AddScores(res []search.SearchResult, query string) []search.SearchResult {
	for i := range res {
		res[i].Score = getGlobalScore(query,res[i].Content)
	}
	sort.Slice(res,func(i,j int) bool { return res[i].Score > res[j].Score })
	return res
}

// func extractContent(reader *bufio.Scanner) string {
// 	var content string
// 	fmt.Println("new content")
// 	for reader.Scan() && reader.Text() != ""  {
// 		fmt.Println("OUT:",reader.Text())
// 		content += reader.Text() + "\n"
// 	}
// 	content = strings.TrimSuffix(content, "\n")
// 	return content	
// }

func GetReadwiseHighlights(txt string) string {
	scan := bufio.NewScanner(strings.NewReader(txt))
	isMetaSkipped := false
	var lines string
	for scan.Scan() {
		line := scan.Text()
		if strings.HasPrefix(line,"## Highlights") {
			isMetaSkipped = true
			scan.Scan() // skip highlight line
			line = scan.Text()
		}
		if isMetaSkipped {
			lines += line + "\n"
		}

	}
	lines = strings.TrimSuffix(lines, "\n")
	return lines
}

func IsReadwiseResult(title string) bool {
	return strings.HasPrefix(title, "Readwise/")
}
