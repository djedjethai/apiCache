package listing

import (
	"errors"
	"fmt"
	"github.com/djedjethai/apiCache/pkg/storage/database"
)

var ErrServer = errors.New("Server error")

type Service interface {
	GetBeersS() ([]Beer, error)
}

type RepoDb interface {
	GetBeers() ([]database.Beer, error)
}

type service struct {
	rdb RepoDb
}

func NewService(rdb RepoDb) service {
	return &service{rdb}
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
