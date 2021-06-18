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
}

type service struct {
	rdb RepoDb
}

func NewService(rdb RepoDb) Service {
	return &service{rdb}
}

func (s *service) BeerDeleteS(id int) error {
	beerFromDB, err := s.rdb.GetBeer(id)
	if err != nil {
		return err
	}

	return s.rdb.BeerDelete(beerFromDB.ID)
}
