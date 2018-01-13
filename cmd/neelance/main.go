package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/nfisher/swapi"
	"github.com/nfisher/swapi/handlers"
)

type RuntimeConfig struct {
	JsonPath string
	Listen   string
}

func main() {
	var config RuntimeConfig

	flag.StringVar(&config.JsonPath, "filename", "swapi.json", "Star Wars JSON API file path.")
	flag.StringVar(&config.Listen, "listen", "127.0.0.1:8008", "HTTP listening address to bind to.")
	flag.Parse()

	_, err := swapi.LoadUniverse(config.JsonPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/graphql", handlers.GQLHandler)
	http.HandleFunc("/", handlers.GraphiQLHandler)

	log.Printf("server listening at %v", config.Listen)
	log.Fatal(http.ListenAndServe(config.Listen, nil))
}
