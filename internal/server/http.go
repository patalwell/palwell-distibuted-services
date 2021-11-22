package server

/*
When building a JSON/HTTP Go server, each handler consists of three steps:
Unmarshal the request’s JSON body into a struct. Run that endpoint’s logic
with the request to obtain a result. Marshal and write that result to the response.

Jeffery, Travis. Distributed Services with Go (p. 26). Pragmatic Bookshelf.

*/

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpServer(addr string) *http.Server {
	httpSrv := newHttpServer()
	r := mux.NewRouter()

	r.HandleFunc("/",httpSrv.handleProduce).Methods("POST")
	r.HandleFunc("/", httpSrv.handleConsume).Methods("GET")
	return &http.Server{
		Addr: addr,
		Handler: r,
	}
}

type httpServer struct {
	Log *Log
}

func newHttpServer() * httpServer {
	return &httpServer{
		Log: NewLog(),
	}
}

type ProduceRequest struct {
	Record Record `json:"record"`
}

type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct {
	Offset uint64 `json:"offset"`
}

type ConsumeResponse struct {
	Record Record `json:"record"`
}

//create Log commit entry
func (s *httpServer) handleProduce(w http.ResponseWriter, r *http.Request) {
	var req ProduceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	off, err := s.Log.Append(req.Record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := ProduceResponse{Offset: off}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//fetch log commit entry
func (s *httpServer) handleConsume(w http.ResponseWriter, r *http.Request) {
	var req ConsumeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	record, err := s.Log.Read(req.Offset)
	if err == fmt.Errorf("offset not found") {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := ConsumeResponse{Record: record}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}