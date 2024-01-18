package main

import (
	"log"
	"net/http"
)

func main() {
	port := "8083"
	if port == "" {
		log.Fatal("PORT env is required")
	}

	instanceID := "mari beajar docker"

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "http method not allowed", http.StatusBadRequest)
			return
		}

		text := "hello world"
		if instanceID != "" {
			text = text + ". from " + instanceID
		}

		w.Write([]byte(text))
	})
	server := new(http.Server)
	server.Handler = mux
	server.Addr = "0.0.0.0:" + port

	log.Println("server starting at", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
