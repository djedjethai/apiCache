package listing

import "time"

// Review defines a beer review
type Review struct {
	ID        int       `json:"id,string"`
	BeerID    int       `json:"beer_id,string"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Score     float32   `json:"score"`
	Text      string    `json:"text"`
	Created   time.Time `json:"created"`
}
