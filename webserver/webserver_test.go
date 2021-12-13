package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"strings"
	"testing"

	"github.com/elchead/misou/config"
	"github.com/elchead/misou/search"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	wsURL := strings.TrimPrefix(url, "http")
	url = fmt.Sprintf("ws%s/api/ws", wsURL)
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	return ws
}

func newSearchRequest(query string) *http.Request {
	parm := url.Values{}
	parm.Add("query", query)

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/search"), nil)
	req.URL.RawQuery = parm.Encode()
	return req
}
func TestEndpoint(t *testing.T) {
	searcher := search.NewSearcher()
	cfg := config.LoadConfig(path.Join("..", "appconfig.json"))
	cfg.InitSources(searcher)
	handler := NewSearchHandler(searcher)
	req := newSearchRequest("naval")
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
	var respObj PayloadResponse
	json.NewDecoder(resp.Body).Decode(&respObj)
	assert.NotEmpty(t, respObj.Data)
	// assert.Equal(t, "readwise", res.Provider)
}

func TestWebsocketEndpoint(t *testing.T) {
	searcher := search.NewSearcher()
	cfg := config.LoadConfig(path.Join("..", "appconfig.json"))
	cfg.InitSources(searcher)
	handler := NewSearchHandler(searcher)
	server := httptest.NewServer(handler)
	conn := mustDialWS(t, server.URL)

	req := PayloadRequest{"hi", "query"}
	err := conn.WriteJSON(&req)
	assert.NoError(t, err)

	var got PayloadResponse
	err = conn.ReadJSON(&got)
	assert.NoError(t, err)
	assert.NotEmpty(t, got.Data)
}
