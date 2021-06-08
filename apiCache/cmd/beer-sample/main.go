package main

import (
	"fmt"
	"github.com/djedjethai/apiCache/pkg/adding"
	"github.com/djedjethai/apiCache/pkg/listing"
	"github.com/djedjethai/apiCache/pkg/storage/database"
	"log"
	"net/http"
)

func main() {
	var adder adding.Service
	var lister listing.Service

	db, err := database.NewStorage()

	adder = adding.NewService(db)
	lister = listing.NewService(db)

	router := rest.Handler(adder, lister)

	fmt.Println("the server is listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
