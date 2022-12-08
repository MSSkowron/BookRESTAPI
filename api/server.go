package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type GoBookAPIServer struct {
	listenAddr string
}

func NewGoBookAPIServer(listenAddr string) *GoBookAPIServer {
	return &GoBookAPIServer{
		listenAddr: listenAddr,
	}
}

func (s *GoBookAPIServer) Run() {
	r := mux.NewRouter()

	r.HandleFunc("/books", s.handleGetBooks).Methods("GET")
	r.HandleFunc("/books", s.handlePostBook).Methods("POST")
	r.HandleFunc("/books/{id}", s.handleGetBookByID).Methods("GET")
	r.HandleFunc("/books/{id}", s.handlePutBookByID).Methods("PUT")
	r.HandleFunc("/books/{id}", s.handleDeleteBookByID).Methods("DELETE")

	log.Println("[GoBookAPI] Server is running on: " + s.listenAddr)

	if err := http.ListenAndServe(s.listenAddr, r); err != nil {
		log.Fatal("[GoBookAPI] Error while running server: " + err.Error())
	}
}

func (s *GoBookAPIServer) handleGetBooks(w http.ResponseWriter, r *http.Request) {

}

func (s *GoBookAPIServer) handlePostBook(w http.ResponseWriter, r *http.Request) {

}

func (s *GoBookAPIServer) handleGetBookByID(w http.ResponseWriter, r *http.Request) {

}

func (s *GoBookAPIServer) handlePutBookByID(w http.ResponseWriter, r *http.Request) {

}

func (s *GoBookAPIServer) handleDeleteBookByID(w http.ResponseWriter, r *http.Request) {

}