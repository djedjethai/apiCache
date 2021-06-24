package updating

import (
	"errors"
	"github.com/djedjethai/apiCache/pkg/storage/database"
	"time"
)

var ErrServer = errors.New("Error during update")
var ErrNotFound = errors.New("Error not found")

type Service interface {
	BeerUpdateS(Beer) error
}

type RepoDb interface {
	BeerUpdate(database.Beer) error
	GetBeer(int) (database.Beer, error)
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

func (s *service) BeerUpdateS(beer Beer) error {
	var b database.Beer

	_, err := s.rdb.GetBeer(beer.ID)
	if err != nil {
		return err
	}

	b.ID = beer.ID
	b.Name = beer.Name
	b.Brewery = beer.Brewery
	b.Abv = beer.Abv
	b.ShortDesc = beer.ShortDesc
	b.Created = time.Now()

	if err := s.rdb.BeerUpdate(b); err != nil {
		return err
	}

	// delete cache
	beersId, _ := s.rdb.GetBeersId()
	_ = s.cch.DeleteCache(beersId)

	return nil
}
