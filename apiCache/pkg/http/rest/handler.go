package rest

import (
	"encoding/json"
	"github.com/djedjethai/apiCache/pkg/adding"
	"github.com/djedjethai/apiCache/pkg/listing"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Handler(a adding.Service, l listing.Service) http.Handler {
	router := httprouter.New()

	router.GET("/beers", GetAllBeersR(l))
	router.POST("/beer", PostBeerR(a))

	return router
}

func GetAllBeersR(l listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		beers, err := l.GetBeersS()
		if err != nil {
			http.Error(w, "Beers unfound", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(beers)
	}
}

func PostBeerR(a adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var newBeer adding.Beer
		if err := decoder.Decode(&newBeer); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		str, err := a.AddBeerS(newBeer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(str)
	}
}
