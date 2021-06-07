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

	for i, bdb := range beersDb {
		beers[i].ID = bdb.ID
		beers[i].Name = bdb.Name
		beers[i].Brewery = bdb.Brewery
		beers[i].Abv = bdb.Abv
		beers[i].ShortDesc = bdb.ShortDesc
		beers[i].Created = bdb.Created
	}

	return beers, nil
}
