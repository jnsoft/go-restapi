package main

import (
	"go-restapi/api-service/handlers"

	"github.com/julienschmidt/httprouter"
)

/*
Define all the routes here.
A new Route entry passed to the routes slice will be automatically
translated to a handler with the NewRouter() function
*/
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

type Routes []Route

func AllRoutes(e handlers.Env) Routes {
	routes := Routes{
		Route{"Index", "GET", "/", handlers.Index},
		Route{"ItemIndex", "GET", "/items", e.ListItems},
		Route{"ItemDetails", "GET", "/items/:id", e.GetItem},
		Route{"ItemDetails", "POST", "/items", e.CreateItem},
	}
	return routes
}
