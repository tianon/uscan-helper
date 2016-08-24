package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/github.com/{user}/{repo}", githubHandler)
	log.Fatal(http.ListenAndServe(":8181", r))
}
