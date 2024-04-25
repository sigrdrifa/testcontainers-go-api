package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type AgeResponse struct {
	Count uint64
	Name  string
	Age   int
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	ageResponse := AgeResponse{
		Count: 1000,
		Name:  name,
		Age:   62,
	}
	json, err := json.Marshal(ageResponse)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	io.WriteString(w, string(json))
}

func main() {
	http.HandleFunc("/", getRoot)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
