package application

import (
	"fmt"
	"net/http"

	"github.com/dindasigma/my-playground/go-es/controllers"
)

func initializeRoutes() {
	router.HandleFunc("/", index).Methods(http.MethodGet)

	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.ItemsController.Get).Methods(http.MethodGet)
	//router.HandleFunc("/items/search", controllers.ItemsController.Search).Methods(http.MethodPost)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "One of My Turns")
}
