package main

import (
	"app/config"
	"app/routes"
	"fmt"
	"log"
	"net/http"
)

func init() {
	router := routes.Router()
	http.Handle("/", router)
}

func main() {
	host := fmt.Sprintf("0.0.0.0:%s", config.Port)
	if config.UseTLS == "" {
		err := http.ListenAndServe(host, nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	} else {
		err := http.ListenAndServeTLS(host, config.CertFile, config.KeyFile, nil)
		if err != nil {
			log.Fatal("ListenAndServeTLS: ", err)
		}
	}

}
