package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"practice/auth/model"
	"practice/auth/repository"
)

func UsersHandler(rep *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			users, err := repository.SelectUsers(rep)
			if err != nil {
				log.Printf("failed to select users: %s", err)
				w.WriteHeader(500)
				return
			}
			resp, err := users.JsonMarshal()
			if err != nil {
				w.WriteHeader(500)
				return
			}
			w.Header().Set("content-type", "application/json")
			w.Write(resp)
			return
		case "POST":
			body, err := io.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(500)
				log.Printf("failed to read request: %s", err)
				return
			}
			var user model.User
			err = json.Unmarshal(body, &user)
			if err != nil {
				w.WriteHeader(500)
				log.Printf("failed to unmarshal request: %s", err)
				return
			}
			err = user.CryptPassword()
			if err != nil {
				if err != nil {
					w.WriteHeader(500)
					log.Printf("failed to crypt password: %s", err)
					return
				}
			}
			err = repository.InsertUser(rep, &user)
			if err != nil {
				w.WriteHeader(500)
				log.Printf("failed to insert users: %s", err)
				return
			}
			resp, err := user.JsonMarshal()
			if err != nil {
				w.WriteHeader(500)
				log.Printf("failed to unmarshal request: %s", err)
				return
			}
			w.Write(resp)
		}
	}
}
