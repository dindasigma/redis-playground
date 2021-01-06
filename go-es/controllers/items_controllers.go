package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/dindasigma/my-playground/go-es/domain/items"
	"github.com/dindasigma/my-playground/go-es/domain/queries"
	"github.com/dindasigma/my-playground/go-es/services"
	"github.com/dindasigma/my-playground/go-es/utils/response"
	"github.com/gorilla/mux"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
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

func (cont *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemsService.Get(itemId)
	if err != nil {
		response.RespondError(w, err)
		return
	}
	response.RespondJson(w, http.StatusOK, item)
}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.RespondError(w, err)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err != nil {
		response.RespondError(w, err)
		return
	}

	items, searchErr := services.ItemsService.Search(query)
	if searchErr != nil {
		response.RespondError(w, searchErr)
		return
	}
	response.RespondJson(w, http.StatusOK, items)
}
