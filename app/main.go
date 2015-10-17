package main

import (
	"log"
	"net/http"
	"app/config"
	"app/routes"
)

func init()  {
	router := routes.Router()
    http.Handle("/", router)
}

func main() {
	if config.UseTSL == "" {
		err := http.ListenAndServe("0.0.0.0:9000", nil)
		if err != nil {
			log.Fatal("ListenAndServe: ", err)
		}
	} else {
		err := http.ListenAndServeTLS("0.0.0.0:9000", config.CertFile, config.KeyFile, nil)
		if err != nil {
			log.Fatal("ListenAndServeTLS: ", err)
		}
	}

}
