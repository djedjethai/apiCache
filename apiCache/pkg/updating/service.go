package updating

import (
	"errors"
	"github.com/djedjethai/apiCache/pkg/storage/database"
	"strconv"
	"time"
)

var ErrServer = errors.New("Error during update")
var ErrNotFound = errors.New("Error not found")

type Service interface {
	BeerUpdateS(int, Beer) error
}

type repoDB interface {
	BeerUpdateR(database.Beer) error
}

type service struct {
	rdb RepoDb
}

func NewService(r RepoDb) Service {
	return &service{r}
}

func (s *service) BeerUpdateS(beer) error {
	var b database.Beer

	b.ID = strconv.Atoi(beer.ID)
	b.Name = beer.Name
	b.Brewery = beer.Brewery
	b.Abv = beer.Abv
	b.ShortDesc = beer.ShortDesc
	b.Created = time.Now()

	if err := s.rdb.BeerUpdate(b); err != nil {
		return err
	}

	return nil
}
