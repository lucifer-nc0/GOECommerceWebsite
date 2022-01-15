package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type User struct {
	Name string
	ID   uuid.UUID
	Type string
}

func (s *Server) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i User
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var inList bool
		index := 0
		for _, v := range s.UsersList {
			if inList = v.Name == i.Name; inList {
				break
			}
			index++
		}
		if !inList {

			i.ID = uuid.New()
			s.UsersList = append(s.UsersList, i)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.UsersList[index]); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
