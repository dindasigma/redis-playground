package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dindasigma/my-playground/go-es/domain/items"
	"github.com/dindasigma/my-playground/go-es/services"
	"github.com/dindasigma/my-playground/go-es/utils/response"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
}

func (cont *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		log.Print(err)
		return
	}

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		response.RespondError(w, createErr)
		return
	}
	response.RespondJson(w, http.StatusCreated, result)
}
