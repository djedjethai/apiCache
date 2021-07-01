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
	GetReviewsS(int) ([]Review, error)
	GetBeerReviewsS(int) (BeerReviews, error)
}

type RepoDb interface {
	GetBeers() ([]database.Beer, error)
	GetBeer(int) (database.Beer, error)
	GetReviews(int) ([]database.Review, error)
}

type Cache interface {
	CacheBeers(cache.Beer) error
	CacheReview(cache.Review) error
	GetCacheBeers() ([]cache.Beer, error)
	GetCacheReviews(int) ([]cache.Review, error)
}

type service struct {
	rdb RepoDb
	cch Cache
}

func NewService(rdb RepoDb, cch Cache) Service {
	return &service{rdb, cch}
}

func (s *service) GetReviewsS(bid int) ([]Review, error) {
	var revs []Review

	// if cache
	revFromCch, _ := s.cch.GetCacheReviews(bid)
	if len(revFromCch) > 0 {
		for i := range revFromCch {
			id, _ := strconv.Atoi(revFromCch[i].ID)
			beerid, _ := strconv.Atoi(revFromCch[i].BeerID)
			revToList := Review{
				ID:        id,
				BeerID:    beerid,
				FirstName: revFromCch[i].FirstName,
				LastName:  revFromCch[i].LastName,
				Score:     revFromCch[i].Score,
				Text:      revFromCch[i].Text,
				Created:   revFromCch[i].Created,
			}

			revs = append(revs, revToList)
		}

		return revs, nil
	}

	// if no cache
	revFromDb, err := s.rdb.GetReviews(bid)
	if err != nil {
		return revs, err
	}

	for i := range revFromDb {
		rev := Review{
			ID:        revFromDb[i].ID,
			BeerID:    revFromDb[i].BeerID,
			FirstName: revFromDb[i].FirstName,
			LastName:  revFromDb[i].LastName,
			Score:     revFromDb[i].Score,
			Text:      revFromDb[i].Text,
			Created:   revFromDb[i].Created,
		}

		revForCch := cache.Review{
			ID:        strconv.Itoa(revFromDb[i].ID),
			BeerID:    strconv.Itoa(revFromDb[i].BeerID),
			FirstName: revFromDb[i].FirstName,
			LastName:  revFromDb[i].LastName,
			Score:     revFromDb[i].Score,
			Text:      revFromDb[i].Text,
			Created:   revFromDb[i].Created,
		}

		_ = s.cch.CacheReview(revForCch)

		revs = append(revs, rev)
	}

	return revs, nil
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

func (s *service) GetBeerReviewsS(id int) (BeerReviews, error) {

	var brvs BeerReviews

	b, err := s.GetBeerS(id)
	if err != nil {
		return brvs, err
	}
	fmt.Printf("beeeer: %v", b)

	var revs []Review
	revs, err = s.GetReviewsS(b.ID)
	if err != nil {
		return brvs, err
	}
	fmt.Printf("reviews: %v", revs)

	brvs = BeerReviews{
		Beer:    b,
		Reviews: revs,
	}

	return brvs, nil
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
