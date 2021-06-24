package deleting

import (
	"github.com/djedjethai/apiCache/pkg/storage/database"
)

type Service interface {
	BeerDeleteS(int) error
}

type RepoDb interface {
	BeerDelete(int) error
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

func (s *service) BeerDeleteS(id int) error {
	beerFromDB, err := s.rdb.GetBeer(id)
	if err != nil {
		return err
	}

	// delete cache
	beersId, _ := s.rdb.GetBeersId()
	_ = s.cch.DeleteCache(beersId)

	return s.rdb.BeerDelete(beerFromDB.ID)
}
