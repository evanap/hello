package main

import (
	"log"
	"net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	helloByte := []byte(`{"message": "Hello, world"}`)
	w.Write(helloByte)
}

func main() {
	s := &server{}
	http.Handle("/hello", s)
	log.Fatal(http.ListenAndServe(":9081", nil))
}
