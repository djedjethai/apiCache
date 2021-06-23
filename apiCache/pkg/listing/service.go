package listing

import (
	// "encoding/json"
	"errors"
	"fmt"
	"github.com/djedjethai/apiCache/pkg/storage/cache"
	"github.com/djedjethai/apiCache/pkg/storage/database"
	"strconv"
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

type Cache interface {
	CacheBeers(cache.Beer) error
	GetCacheBeers() ([]cache.Beer, error)
}

type service struct {
	rdb RepoDb
	cch Cache
}

func NewService(rdb RepoDb, cch Cache) Service {
	return &service{rdb, cch}
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

	// get beers from cache
	listBeers, _ := s.cch.GetCacheBeers()

	fmt.Printf("nbr from cache: %v", len(listBeers))
	if len(listBeers) > 0 {
		fmt.Println("from cache")
		for _, beer := range listBeers {
			id, _ := strconv.Atoi(beer.ID)
			b := Beer{
				ID:        id,
				Name:      beer.Name,
				Brewery:   beer.Brewery,
				Abv:       beer.Abv,
				ShortDesc: beer.ShortDesc,
				Created:   beer.Created,
			}

			beers = append(beers, b)
		}

		return beers, nil
	}

	// if no cache
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

		id := strconv.Itoa(b.ID)
		bcch := cache.Beer{
			ID:        id,
			Name:      b.Name,
			Brewery:   b.Brewery,
			Abv:       b.Abv,
			ShortDesc: b.ShortDesc,
			Created:   b.Created,
		}

		_ = s.cch.CacheBeers(bcch)
	}

	return beers, nil
}
