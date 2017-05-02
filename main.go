package main

import (
	"log"
	"net/http"
	"tantan-demo/router"
	"tantan-demo/util"
)

func main() {
	util.InitDbs()

	router := router.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
