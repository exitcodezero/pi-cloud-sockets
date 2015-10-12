package main

import (
	"log"
	"net/http"
	"app/routes"
)

func init()  {
	router := routes.Router()
    http.Handle("/", router)
}

func main() {
	err := http.ListenAndServe("0.0.0.0:9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
