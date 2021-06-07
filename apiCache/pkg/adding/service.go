package adding

import (
	"errors"
	"fmt"
	"github.com/djedjethai/apiCache/pkg/storage"
	"github.com/djedjethai/apiCache/pkg/storage/database"
)

var ErrDuplicate = errors.New("Beer already exist")
var ErrInput = errors.New("Incorrect input")
var ErrRegister = errors.New("Server error")

const (
	beerC = "beer"
)

type Service interface {
	AddBeerS(Beer) (string, error)
	AddBeerSampleS([]Beer) error
}

type RepoDb interface {
	AddBeer(database.Beer) (string, error)
}

type service struct {
	rdb RepoDb
}

func NewService(rdb RepoDb) service {
	return &service{rdb}
}

func (s *service) AddBeerS(beer Beer) (string, error) {
	var beerForDb database.Beer

	// make the id, a bullshit one could be....
	id, err := storage.GetId(beerC)
	if err != nil {
		return "", err
	}

	beerForDb.ID = id
	beerForDb.Name = beer.Name
	beerForDb.Brewery = beer.Brewery
	beerForDb.Abv = beer.Abv
	beerForDb.ShortDesc = beer.ShortDesc
	beerForDb.Created = time.Now()

	// question is how to get the id back from db in a single req
	str, err := s.sdb.AddBeer(beerForDb)
	if err != nil {
		return "", err
	}

	// delete cache

	return str, nil

}

func (s *service) AddBeerSampleS(beers []Beer) error {
	for i := range beers {
		if i == nil {
			return ErrInput
		}

		_, err := s.AddBeerS(beers[i])
		if err != nil {
			return err
		}
	}

	// delete cache
	return nil
}
