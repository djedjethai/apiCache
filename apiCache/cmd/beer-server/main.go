package main

import (
	"fmt"
	"github.com/djedjethai/apiCache/pkg/adding"
	"github.com/djedjethai/apiCache/pkg/deleting"
	"github.com/djedjethai/apiCache/pkg/http/rest"
	"github.com/djedjethai/apiCache/pkg/listing"
	"github.com/djedjethai/apiCache/pkg/reviewing"
	"github.com/djedjethai/apiCache/pkg/storage/cache"
	"github.com/djedjethai/apiCache/pkg/storage/database"
	"github.com/djedjethai/apiCache/pkg/updating"
	"log"
	"net/http"
)

func main() {
	var adder adding.Service
	var lister listing.Service
	var updater updating.Service
	var deleter deleting.Service
	var reviewer reviewing.Service

	db, _ := database.NewStorage()
	cch, _ := cache.NewStorage()

	adder = adding.NewService(db, cch)
	lister = listing.NewService(db, cch)
	updater = updating.NewService(db, cch)
	deleter = deleting.NewService(db, cch)
	reviewer = reviewing.NewService(db)

	router := rest.Handler(adder, lister, updater, deleter, reviewer)

	fmt.Println("the server is listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
