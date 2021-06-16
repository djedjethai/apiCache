package main

import (
	"fmt"
	"github.com/djedjethai/apiCache/pkg/adding"
	"github.com/djedjethai/apiCache/pkg/http/rest"
	"github.com/djedjethai/apiCache/pkg/listing"
	"github.com/djedjethai/apiCache/pkg/storage/database"
	"github.com/djedjethai/apiCache/pkg/updating"
	"log"
	"net/http"
)

func main() {
	var adder adding.Service
	var lister listing.Service
	var updater updating.Service

	db, _ := database.NewStorage()

	adder = adding.NewService(db)
	lister = listing.NewService(db)
	updater = updating.NewService(db)

	router := rest.Handler(adder, lister, updater)

	fmt.Println("the server is listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
