package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"practice/auth/model"
	"practice/auth/repository"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func AuthHandler(rep *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			auth, err := requestToJson(r)
			if err != nil {
				w.WriteHeader(500)
				log.Printf("failed to login: %s", err)
				return
			}
			user, err := repository.SearchForEmail(rep, auth.Email)
			if err != nil || user == nil {
				w.WriteHeader(500)
				log.Printf("failed to login: %s", err)
				return
			}
			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(auth.Password))
			if err != nil {
				w.WriteHeader(500)
				log.Printf("failed to login: %s", err)
				return
			}
			claims := jwt.StandardClaims{
				Issuer:    strconv.Itoa(int(user.ID)),
				ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			}
			jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			token, err := jwtToken.SignedString([]byte("secret"))
			if err != nil {
				w.WriteHeader(500)
				log.Printf("failed to login: %s", err)
				return
			}
			cokkie := http.Cookie{
				Name:     "jwt",
				Value:    token,
				Expires:  time.Now().Add(time.Hour * 24),
				HttpOnly: true,
			}
			http.SetCookie(w, &cokkie)
		}
	}
}

func requestToJson(r *http.Request) (*model.AuthMessage, error) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}

	var auth model.AuthMessage
	err = json.Unmarshal(b, &auth)
	if err != nil {
		return nil, err
	}
	return &auth, nil
}
