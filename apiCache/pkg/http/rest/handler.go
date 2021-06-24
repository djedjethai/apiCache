package rest

import (
	"encoding/json"
	"github.com/djedjethai/apiCache/pkg/adding"
	"github.com/djedjethai/apiCache/pkg/deleting"
	"github.com/djedjethai/apiCache/pkg/listing"
	"github.com/djedjethai/apiCache/pkg/updating"
	"github.com/djedjethai/apiCache/pkg/reviewing"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func Handler(a adding.Service, l listing.Service, u updating.Service, d deleting.Service, r reviewing.Service) http.Handler {
	router := httprouter.New()

	router.GET("/beers", GetAllBeersR(l))
	router.GET("/beer/:id", GetBeerR(l))
	router.POST("/beer", PostBeerR(a))
	router.POST("/beer/update", PostBeerUpdateR(u))
	router.DELETE("/beer/:id", DeleteBeerR(d))
	router.POST("/review/:id", PostReview(r))

	return router
}

func PostReview(r reviewing.Service) func(...)

func DeleteBeerR(d deleting.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		beerId, err := strconv.Atoi(p.ByName("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := d.BeerDeleteS(beerId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Beer deleted")
	}
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

func PostBeerUpdateR(u updating.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var beer updating.Beer
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&beer); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if err := u.BeerUpdateS(beer); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("Beer updated")
	}
}
