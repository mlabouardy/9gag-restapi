package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlabouardy/9gag"
)

// Get list of memes by their tag
func MemesByTagEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gag9 := gag9.New()
	memes := gag9.FindByTag(params["tag"])
	json.NewEncoder(w).Encode(memes)
}

// Expose a new route and start an HTTP server on port 3000
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/memes/{tag}", MemesByTagEndpoint)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
