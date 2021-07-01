package listing

type BeerReviews struct {
	Beer    Beer     `json:"beer"`
	Reviews []Review `json:"reviews"`
}
