package main

import (
	"github.com/exitcodezero/picloud/config"
	"github.com/exitcodezero/picloud/routes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func init() {
	router := routes.Router()
	http.Handle("/", router)
}

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

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
