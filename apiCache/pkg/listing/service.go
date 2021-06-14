package listing

import (
	"errors"
	// "fmt"
	"github.com/djedjethai/apiCache/pkg/storage/database"
)

var ErrServer = errors.New("Server error")

type Service interface {
	GetBeersS() ([]Beer, error)
	GetBeerS(int) (Beer, error)
}

type RepoDb interface {
	GetBeers() ([]database.Beer, error)
	GetBeer(int) (database.Beer, error)
}

type service struct {
	rdb RepoDb
}

func NewService(rdb RepoDb) Service {
	return &service{rdb}
}

func (s *service) GetBeerS(id int) (Beer, error) {
	var b Beer

	beerFromDB, err := s.rdb.GetBeer(id)
	if err != nil {
		return b, err
	}

	b.ID = beerFromDB.ID
	b.Name = beerFromDB.Name
	b.Brewery = beerFromDB.Brewery
	b.Abv = beerFromDB.Abv
	b.ShortDesc = beerFromDB.ShortDesc
	b.Created = beerFromDB.Created

	return b, nil
}

func (s *service) GetBeersS() ([]Beer, error) {
	var beers []Beer

	beersDb, err := s.rdb.GetBeers()
	if err != nil {
		return beers, err
	}

	for i := range beersDb {
		b := Beer{
			ID:        beersDb[i].ID,
			Name:      beersDb[i].Name,
			Brewery:   beersDb[i].Brewery,
			Abv:       beersDb[i].Abv,
			ShortDesc: beersDb[i].ShortDesc,
			Created:   beersDb[i].Created,
		}

		beers = append(beers, b)
	}

	return beers, nil
}
