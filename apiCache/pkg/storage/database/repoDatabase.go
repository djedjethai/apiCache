package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "root"
	hostname = "mysql:3306"
	database = "apiCache"
)

type Storage struct {
	db *sql.DB
}

func dsn() string {
	return fmt.Sprintf("%s:%s@TCP(%s)/%s?charset=utf8", username, password, hostname, database)
}

func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	s.db, err = sql.Open("mysql", dsn())
	fmt.Println(err)
	defer s.db.Close()

	err = s.db.Ping()
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	query := `CREATE TABLE IF NOT EXISTS beer(beer_id int primary key auto_increment, beer_name VARCHAR(20), beer_brewery VARCHAR(20), beer_abv FLOAT(25), beer_shortdesc text, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}

	r, err := stmt.Exec()
	fmt.Println(err)

	n, err := r.RowsAffected()
	fmt.Println(err)

	fmt.Printf("table beer created: %v", n)

	return s, nil
}

func (s *Storage) AddBeer(b Beer) (string, error) {
	fmt.Println("AddBeer func in db triggered, coool")

	return "", nil
}

func (s *Storage) GetBeers() ([]Beer, error) {
	var beers []Beer
	fmt.Println("GetBeers func in db triggered, ras")

	return beers, nil
}
