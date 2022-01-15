package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct {
	ID       uuid.UUID
	Name     string
	Quantity int
	Price    int
}

type Server struct {
	*mux.Router
	itemsList []Item
	UsersList []User
}

func NewServer() *Server {
	s := &Server{
		Router:    mux.NewRouter(),
		itemsList: []Item{},
		UsersList: []User{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/items", s.createItem()).Methods("POST")
	s.HandleFunc("/register", s.CreateUser()).Methods("POST")
	s.HandleFunc("/items", s.listItems()).Methods("GET")
	s.HandleFunc("/buy/{product}", s.buyItem()).Methods("GET")
}

func (s *Server) listItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.itemsList); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func (s *Server) buyItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		index := 0
		for _, v := range s.itemsList {
			if v.Name == vars["product"] {
				if v.Quantity == 0 {
					break
				}
				s.itemsList[index].Quantity--
				break
			}
			index++
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.itemsList[index]); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func (s *Server) createItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var inList bool
		index := 0
		for _, v := range s.itemsList {
			if inList = v.Name == i.Name; inList {
				s.itemsList[index].Quantity++
				break
			}
			index++
		}
		if !inList {
			i.Quantity = 1
			i.ID = uuid.New()

			s.itemsList = append(s.itemsList, i)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.itemsList[index]); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
