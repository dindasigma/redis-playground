package application

import (
	"log"
	"net/http"
	"time"

	"github.com/dindasigma/my-playground/go-es/client/elasticsearch"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()

	initializeRoutes()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Print("about to start the application on port 8000...")
	log.Fatal(srv.ListenAndServe())
}
