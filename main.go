package main

import (
	"log"
	"net/http"

	"github.com/xandeer/alpha-api/handlers"
	"github.com/xandeer/alpha-api/routers"
)

// Define HTTP request routes
func main() {
	router := routers.NewRouter()

	if err := handlers.InitOperateVersion(); err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
