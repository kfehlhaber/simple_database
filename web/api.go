package web

import (
	"fmt"
	"github.com/kejne/simple_database/db"
	"net/http"
)

func Serve() {
	http.HandleFunc("/get", get)
	http.HandleFunc("/set", set)
	fmt.Println("Web API server ready to accept requests")
	http.ListenAndServe(":80", nil)
}

func get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")

	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		response := db.Fetch(key)

		if response == "" {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(response))
		}
	}
}

func set(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	for key := range query {
		response := []byte(db.Persist(key, query.Get(key)))
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		// Ignore all but first parameter
		break;
	}
}
