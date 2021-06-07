package database

import(
	"fmt"
)

type RepoDb struct{
	GetBeers()([]Beer, error)
	AddBeer([]Beer, error)
} 
