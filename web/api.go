package web

import (
	"github.com/kejne/simple_database/db"
	"log"
	"net/http"
)

func Serve() {
	http.HandleFunc("/get", get)
	http.HandleFunc("/set", set)
	log.Println("Web API server ready to accept requests")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
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
			n, err := w.Write([]byte(response))

			if n < len(response) {
				log.Fatalln(err.Error())
			}
		}
	}
}

func set(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	for key := range query {
		response := []byte(db.Persist(key, query.Get(key)))
		w.WriteHeader(http.StatusOK)
		n, err := w.Write(response)
		if n < len(response) {
			log.Fatalln(err.Error())
		}
		// Ignore all but first parameter
		break;
	}
}
