package cache

import (
	"encoding/json"
	// "fmt"
	"github.com/djedjethai/apiCache/pkg/storage/database"
	"github.com/nanobox-io/golang-scribble"
	"path"
	"runtime"
	"strconv"
)

const (
	dir = "/data/"

	CollectionBeer   = "beers"
	CollectionReview = "review"
)

type RepoDb interface {
	GetBeersID() ([]database.Beer, error)
}

type Storage struct {
	cch *scribble.Driver
}

func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.cch, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}

	return s, err
}

func (s *Storage) CacheBeers(b Beer) error {
	if err := s.cch.Write(CollectionBeer, b.ID, b); err != nil {
		return err
	}

	return nil
}

func (s *Storage) CacheReview(r Review) error {
	if err := s.cch.Write(CollectionReview, r.ID, r); err != nil {
		return err
	}

	return nil
}

func (s *Storage) GetCacheBeer(bid int) (Beer, error) {
	var beer Beer
	id := strconv.Itoa(bid)
	if err := s.cch.Read(CollectionBeer, id, &beer); err != nil {
		return beer, err
	}

	return beer, nil
}

func (s *Storage) GetCacheReviews(bid int) ([]Review, error) {
	var revs []Review
	listReviews, err := s.cch.ReadAll(CollectionReview)
	if err != nil {
		return revs, err
	}

	// parse the reviews and match b id
	beer_id := strconv.Itoa(bid)
	for _, rev := range listReviews {
		var review Review
		_ = json.Unmarshal([]byte(rev), &review)

		if review.BeerID == beer_id {
			revs = append(revs, review)
		}
	}

	return revs, nil
}

func (s *Storage) GetCacheBeers() ([]Beer, error) {
	var listBeers []Beer

	listBeerCch, err := s.cch.ReadAll(CollectionBeer)
	if err != nil {
		return listBeers, err
	}

	for _, brBin := range listBeerCch {
		var beer Beer
		_ = json.Unmarshal([]byte(brBin), &beer)

		listBeers = append(listBeers, beer)
	}

	return listBeers, nil
}

func (s *Storage) DeleteCache(ids []database.Beer) error {
	for i := range ids {
		// delete each entry in cache
		idstr := strconv.Itoa(ids[i].ID)
		if err := s.cch.Delete(CollectionBeer, idstr); err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) DeleteCacheReview(revs []database.Review) error {
	// itere and delete each
	for i := range revs {
		bid := strconv.Itoa(revs[i].BeerID)
		if err := s.cch.Delete(CollectionReview, bid); err != nil {
			return err
		}
	}

	return nil
}
