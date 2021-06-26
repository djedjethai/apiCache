package reviewing

// import "encoding/json"

// Review defines a beer review
type Review struct {
	BeerID    int     `json:"beer_id,string"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Score     float32 `json:"score"`
	Text      string  `json:"text"`
}
