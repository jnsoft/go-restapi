package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	*mux.Router
	items []Item
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
		items:  []Item{},
	}
	s.routes()
	s.initDb(25)
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/items", s.listItems()).Methods("GET")
	s.HandleFunc("/items/", s.listItems()).Methods("GET")
	s.HandleFunc("/items/{id}", s.getItem()).Methods("GET")
	s.HandleFunc("/items", s.createItem()).Methods("POST")
	s.HandleFunc("/items/{id}", s.removeItem()).Methods("DELETE")
}

func (s *Server) initDb(n int) {

	for i := 0; i < n; i++ {
		s.items = append(s.items, Item{Id: uuid.New(), Name: fmt.Sprintf("Item %v", i+1)})
	}
}

func (s *Server) listItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.items); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) getItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id_str := mux.Vars(r)["id"]
		id, err := uuid.Parse(id_str)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for _, item := range s.items {
			if item.Id == id {
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(item); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}
		}
		http.NotFound(w, r)
	}
}

func (s *Server) createItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		i.Id = uuid.New()
		s.items = append(s.items, i)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) removeItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id_str := mux.Vars(r)["id"]
		id, err := uuid.Parse(id_str)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		for i, item := range s.items {
			if item.Id == id {
				s.items = append(s.items[:i], s.items[i+1:]...)
				return
			}
		}

		http.NotFound(w, r)
	}
}
