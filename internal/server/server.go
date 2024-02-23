package server

import (
	"net/http"

	"github.com/go-chi/chi"

	log "github.com/sirupsen/logrus"
)

func Init() {
	const PORT = "4001"
	router := chi.NewRouter()

	// Init mongo

	log.SetFormatter(&log.JSONFormatter{})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
