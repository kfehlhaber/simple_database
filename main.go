package main

import (
	"net/http"
)

var dbMap map[string]string

func main() {
	dbMap = make(map[string]string)
	http.HandleFunc("/get", get)
	http.HandleFunc("/set", set)

	http.ListenAndServe(":80", nil)
}

func get(w http.ResponseWriter, r *http.Request) {
	var response = []byte(dbMap[r.URL.Query().Get("key")])
	w.Write(response)
}

func set(w http.ResponseWriter, r *http.Request) {
	for key := range r.URL.Query() {
		dbMap[key] = r.URL.Query().Get(key)
	}
}
