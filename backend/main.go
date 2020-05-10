package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gothub.com/piyusgupta/pgtuner/backend/api"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
)

func main() {
	flag.Parse()
	r := httprouter.New()
	r.GET("/db/settings/", api.AllPGSettings)
	// r.GET("/db/tables/", api.pgtables.List)
	err := http.ListenAndServe(*addr, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
