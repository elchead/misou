package webserver

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/elchead/misou/search"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type SearchServer struct {
	searcher *search.Searcher
	http.Handler
}

type PayloadRequest struct {
	Query  string `json:"query"`
	Action string `json:"action"`
}

type PayloadResponse struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"` //[]source.SearchResult `json:"data"`
	Query  string      `json:"query"`
}

func readWsRequest(conn *websocket.Conn) (*PayloadRequest, error) {
	req := &PayloadRequest{}
	err := conn.ReadJSON(req)
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Fatalf("unexpected close error: %v\n", err)
		}
		return nil, err
	}
	return req, nil
}

func (s *SearchServer) searchHandler(w http.ResponseWriter, r *http.Request) {
	query, ok := r.URL.Query()["query"]
	if !ok {
		log.Info("no query")
	}

	res, err := s.searcher.Search(query[0])
	if err != nil {
		log.Warningf("Search %s failed: %v\n", query, err)
	}
	payload := PayloadResponse{Action: "", Data: res, Query: query[0]}
	json.NewEncoder(w).Encode(payload)
}

const (
	ACTION_RESULTS        = "results"
	ACTION_LOADING_STATUS = "loading_status"
)

type LoadingStatus struct {
	Provider string `json:"provider"`
	Loading  bool   `json:"loading"`
}

func (s *SearchServer) sendSearchResults(conn *websocket.Conn, query string, writer chan PayloadResponse) {
	go func() {
		writer <- PayloadResponse{Action: ACTION_LOADING_STATUS, Data: LoadingStatus{Loading: true, Provider: "hans"}, Query: query}
		res, err := s.searcher.Search(query)
		if err != nil {
			log.Infof("Search error: %v", err)
		} else {
			results_resp := PayloadResponse{Action: ACTION_RESULTS, Data: res, Query: query}
			writer <- results_resp
		}
		writer <- PayloadResponse{Action: ACTION_LOADING_STATUS, Data: LoadingStatus{Loading: false, Provider: ""}, Query: query}

	}()
	// for _, client := range s.SearchClients {
	// }
}

func (s *SearchServer) readRequest(conn *websocket.Conn) (*PayloadRequest, error) {
	req := &PayloadRequest{}
	err := conn.ReadJSON(req)
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Fatalf("unexpected close error: %v\n", err)
		}
		return nil, err
	}
	return req, nil
}

func (s *SearchServer) wsSearchHandler(w http.ResponseWriter, r *http.Request) {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}
	writer := make(chan PayloadResponse, 5)
	done := make(chan bool, 1)

	go func() {
		for {
			select {
			case resp := <-writer:
				err = conn.WriteJSON(&resp)
				if err != nil {
					log.Printf("Could not write JSON (%+v), result: %v", resp, err)
					return
				}
			case <-done:
				return
			}
		}
	}()

	for {
		req, err := s.readRequest(conn)
		if err != nil {
			break
		}
		log.Printf("Request: %s\n", req.Query)
		s.sendSearchResults(conn, req.Query, writer)
	}
	done <- true

	// my
	// req, _ := readWsRequest(conn)
	// query := req.Query

	// // do search
	// res, err := s.searcher.Search(query)
	// if err != nil {
	// 	log.Fatalf("Search error: %v", err)
	// }

	// // write response
	// payload := PayloadResponse{Action: "", Data: res, Query: query}
	// err = conn.WriteJSON(&payload)
	// if err != nil {
	// 	log.Printf("Could not write JSON (%+v), result: %v", payload, err)
	// 	return
	// }

	defer conn.Close()
}

func NewSearchHandler(searcher *search.Searcher) http.Handler {
	server := SearchServer{searcher: searcher}

	router := http.NewServeMux()
	router.Handle("/search", http.HandlerFunc(server.searchHandler))
	router.Handle("/api/ws", http.HandlerFunc(server.wsSearchHandler))
	return router
}
