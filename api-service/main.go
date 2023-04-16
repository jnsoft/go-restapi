package main

import (
	"go-restapi/api-service/handlers"
	"go-restapi/api-service/store"
	"log"
	"net/http"
)

func main() {
	store := &store.ItemStore{}
	store.Init(5)
	env := handlers.Env{Db: store}

	router := NewRouter(AllRoutes(env))
	log.Fatal(http.ListenAndServe(":8082", router))
}
