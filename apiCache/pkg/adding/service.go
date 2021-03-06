package adding

import (
	"errors"
	// "fmt"
	"github.com/djedjethai/apiCache/pkg/storage/database"
	"time"
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
	GetBeersId() ([]database.Beer, error)
}

type Cache interface {
	DeleteCache([]database.Beer) error
}

type service struct {
	rdb RepoDb
	cch Cache
}

func NewService(rdb RepoDb, cch Cache) Service {
	return &service{rdb, cch}
}

func (s *service) AddBeerS(beer Beer) (string, error) {
	var beerForDb database.Beer

	// what ever id, it won't be send to db
	beerForDb.ID = 0
	beerForDb.Name = beer.Name
	beerForDb.Brewery = beer.Brewery
	beerForDb.Abv = beer.Abv
	beerForDb.ShortDesc = beer.ShortDesc
	beerForDb.Created = time.Now()

	// question is how to get the id back from db in a single req
	str, err := s.rdb.AddBeer(beerForDb)
	if err != nil {
		return "", err
	}

	// delete cache
	bids, _ := s.rdb.GetBeersId()
	_ = s.cch.DeleteCache(bids)

	return str, nil
}

func (s *service) AddBeerSampleS(beers []Beer) error {
	for i := range beers {
		_, err := s.AddBeerS(beers[i])
		if err != nil {
			return err
		}
	}

	return nil
}
