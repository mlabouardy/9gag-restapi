package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlabouardy/9gag"
)

func MemesByTagEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gag9 := gag9.New()
	memes := gag9.FindByTag(params["tag"])
	json.NewEncoder(w).Encode(memes)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/memes/{tag}", MemesByTagEndpoint)
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
