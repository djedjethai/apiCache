package reviewing

import (
	"github.com/djedjethai/apiCache/pkg/storage/database"
)

type RepoDb interface {
	AddReview(database.Review) (string, error)
}

type Service interface {
	AddReviewS(Review) (string, error)
}

type service struct {
	rdb RepoDb
}

func NewService(rdb RepoDb) Service {
	return &service{rdb}
}

func (s *service) AddReviewS(r Review) (string, error) {
	var revForDb database.Review

	revForDb.ID = 0
	revForDb.BeerID = r.BeerID
	revForDb.FirstName = r.FirstName
	revForDb.LastName = r.LastName
	revForDb.Score = r.Score
	revForDb.Text = r.Text

	str, err := s.rdb.AddReview(revForDb)
	if err != nil {
		return "", err
	}

	return str, nil
}
