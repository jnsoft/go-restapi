package api

import "github.com/gorilla/mux"

type Server struct {
	*mux.Router
	items []Item
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		items:  []Item{},
	}
	return s
}
