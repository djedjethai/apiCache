package rest

import (
	"encoding/json"
	"github.com/djedjethai/apiCache/pkg/adding"
	"github.com/djedjethai/apiCache/pkg/listing"
	"github.com/djedjethai/apiCache/pkg/updating"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func Handler(a adding.Service, l listing.Service, u updating.Service) http.Handler {
	router := httprouter.New()

	router.GET("/beers", GetAllBeersR(l))
	router.GET("/beer/:id", GetBeerR(l))
	router.POST("/beer", PostBeerR(a))
	router.POST("/:id/update", PostBeerUpdateR(u))

	return router
}

func GetBeerR(l listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		beerParam, _ := strconv.Atoi(p.ByName("id"))

		beer, err := l.GetBeerS(beerParam)
		if err != nil {
			http.Error(w, "Beer unfound", http.StatusNotFound)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(beer)
	}
}

func GetAllBeersR(l listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		beers, err := l.GetBeersS()
		if err != nil {
			http.Error(w, "Beers unfound", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
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

func PostBeerUpdateR(a adding.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	}
}
