package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/piyusgupta/pgtuner/backend/api"
)

var (
	addr = flag.String("addr", ":8080", "http service address")
)

func main() {
	flag.Parse()
	router := mux.NewRouter()
	router.HandleFunc("/db/settings/", api.PGSettingHandler)
	log.Println("started server  at port", *addr)
	err := http.ListenAndServe(*addr, router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
