package main

import (
	"net/http"
	"practice/auth/repository"
	"practice/auth/router"
)

func main() {
	rep := repository.NewRepository()
	router.SetRouter(rep)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
