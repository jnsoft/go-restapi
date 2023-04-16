package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"go-restapi/api-service/models"
	"go-restapi/api-service/store"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type Env struct {
	Db *store.ItemStore
}

// GET /items
func (e *Env) ListItems(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	items := []*models.Item{}
	for _, item := range *e.Db {
		items = append(items, item)
	}
	JsonResponse(w, items)
}

// GET /items/:id
func (e *Env) GetItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id_str := params.ByName("id")
	id, err := uuid.Parse(id_str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item, ok := e.Db.Find(id)
	if !ok {
		writeErrorResponse(w, http.StatusNotFound, "Item Not Found")
		return
	}
	JsonResponse(w, item)
}

// POST /items
func (e *Env) CreateItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	item := &models.Item{}
	if err := populateModelFromHandler(w, r, params, item); err != nil {
		writeErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	item_store[book.ISDN] = book
	writeOKResponse(w, book)
}

// Writes the response as a standard JSON response with StatusOK
func JsonResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(m); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Writes the response as a standard JSON response with StatusOK
func writeOKResponse_meta(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&models.JsonResponse{Data: m}); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// Writes the error response as a Standard API JSON response with a response code
func writeErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.
		NewEncoder(w).
		Encode(&models.JsonErrorResponse{Error: &models.ApiError{Status: errorCode, Title: errorMsg}})
}

// Populates a model from the params in the Handler
func populateModelFromHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params, model interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}
