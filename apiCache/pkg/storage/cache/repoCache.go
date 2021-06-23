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

	CollectionBeer = "beers"
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
