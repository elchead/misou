package source

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/elchead/misou/search"
	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

const gdriveProvider = "gdrive"

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config, tokenReader io.Reader, tokenWriter io.Writer) (*http.Client, *oauth2.Token) {
	tok, err := getTokenFromFile(tokenReader)
	// TODO check token validity
	if err != nil {
		log.Infof("Need to update Gdrive token: %s", err)
		tok = getTokenFromWeb(config)
		saveToken(tokenWriter, tok)
	}
	return config.Client(context.Background(), tok), tok
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	reader := bufio.NewReader(os.Stdin)

	authCode, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func getTokenFromFile(f io.Reader) (*oauth2.Token, error) {
	tok := &oauth2.Token{}
	err := json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(f io.Writer, token *oauth2.Token) {
	token.AccessToken = "hi"
	err := json.NewEncoder(f).Encode(token)
	if err != nil {
		log.Fatal(err)
	}
}

type GoogleDriveClient struct {
	drive *drive.Service
	token *oauth2.Token
}

func NewGdriveClient(path string, credentialReader io.Reader, tokenReader io.Reader, tokenWriter io.Writer) GoogleDriveClient { //DataIntegrator
	credentials, err := ioutil.ReadAll(credentialReader)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	ctx := context.Background()
	config, err := google.ConfigFromJSON(credentials, drive.DriveScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client, token := getClient(config, tokenReader, tokenWriter)
	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}
	return GoogleDriveClient{srv, token}
}

func (client GoogleDriveClient) Search(keywords string) ([]search.SearchResult, error) {
	resp, err := client.getRawSearchResponse(keywords)
	if err != nil {
		return nil, errors.Wrap(err, "Google drive search query failed")
	}
	return client.parseFiles(resp.Files), nil
}

func (client *GoogleDriveClient) getRawSearchResponse(keywords string) (*drive.FileList, error) {
	query := fmt.Sprintf("fullText contains '%s'", keywords)
	return client.drive.Files.List().Q(query).Fields("files(*)").Do()
}

func (client GoogleDriveClient) parseFiles(files []*drive.File) []search.SearchResult {
	res := []search.SearchResult{}
	for _, file := range files {
		// TODO add content through Download of file?
		data := search.SearchResult{Title: file.Name, Link: file.WebViewLink, Content: "", Provider: gdriveProvider, ContentType: file.FileExtension}
		res = append(res, data)
	}
	return res
}
