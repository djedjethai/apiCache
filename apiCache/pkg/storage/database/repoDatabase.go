package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"time"
)

// const (
// 	username = "root"
// 	password = "root"
// 	hostname = "mysql:3306"
// 	database = "apiCache"
// )

type Storage struct {
	db *sql.DB
}

// func dsn() string {
// 	return fmt.Sprintf("%s:%s@TCP(%s)/%s?parseTime=true", username, password, hostname, database)
// }

func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	// s.db, err = sql.Open("mysql", dsn())
	s.db, err = sql.Open("mysql", "root:root@tcp(mysql:3306)/apiCache?parseTime=true")

	fmt.Println(err)
	// defer s.db.Close() // if on, close the connection

	err = s.db.Ping()
	fmt.Printf("grrr: %v", err)
	if err != nil {
		return nil, err
	}

	query := `CREATE TABLE IF NOT EXISTS beer(beer_id int primary key auto_increment, beer_name VARCHAR(20), beer_brewery VARCHAR(20), beer_abv FLOAT(25), beer_shortdesc text, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	res, err := s.db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when creating product table", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected", err)
	}

	fmt.Printf("table beer created: %v", rows)

	return s, nil
}

func (s *Storage) AddBeer(b Beer) (string, error) {

	query := "INSERT INTO beer(beer_name, beer_brewery, beer_abv, beer_shortdesc) VALUES (?, ?, ?, ?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	fmt.Println("state1")

	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return "", err
	}
	defer stmt.Close()

	fmt.Println("bf exec the stmt")

	res, err := stmt.ExecContext(ctx, b.Name, b.Brewery, b.Abv, b.ShortDesc)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
		return "", err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return "", err
	}

	// get last row id insered (in this case the row id)
	prdID, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error %s when getting last inserted product", err)
		return "", err
	}

	log.Printf("%d products created ", rows)

	// prdID is int64, so convert to int then to str
	return strconv.Itoa(int(prdID)), nil
}

func (s *Storage) GetBeers() ([]Beer, error) {
	var beers []Beer
	fmt.Println("GetBeers func in db triggered, ras")

	results, err := s.db.Query("SELECT beer_id, beer_name, beer_brewery, beer_abv, beer_shortdesc, created_at FROM beer")
	if err != nil {
		// fmt.Println("allllooo")
		return beers, err
	}

	fmt.Printf("beers bf scan %v", results)
	for results.Next() {
		var b Beer
		err := results.Scan(
			&b.ID,
			&b.Name,
			&b.Brewery,
			&b.Abv,
			&b.ShortDesc,
			&b.Created,
		)
		fmt.Printf("beeer: %v", b)
		if err != nil {
			return beers, err
		}

		beers = append(beers, b)
	}

	return beers, nil
}

func (s *Storage) GetBeersId() ([]Beer, error) {
	var beersId []Beer

	results, err := s.db.Query("SELECT beer_id FROM beer")
	if err != nil {
		// fmt.Println("allllooo")
		return beersId, err
	}

	for results.Next() {
		var b Beer
		err := results.Scan(
			&b.ID,
		)
		if err != nil {
			return beersId, err
		}

		beersId = append(beersId, b)
	}

	return beersId, nil
}

func (s *Storage) GetBeer(bid int) (Beer, error) {
	var b Beer

	err := s.db.QueryRow("SELECT beer_id, beer_name, beer_brewery, beer_abv, beer_shortdesc, created_at FROM beer WHERE beer_id = ?", bid).Scan(
		&b.ID,
		&b.Name,
		&b.Brewery,
		&b.Abv,
		&b.ShortDesc,
		&b.Created,
	)
	if err != nil {
		fmt.Printf("grrr: %v", err)
		return b, err
	}

	return b, nil

}

func (s Storage) BeerUpdate(b Beer) error {

	stmt, err := s.db.Prepare("UPDATE beer SET beer_name=?, beer_brewery=?, beer_abv=?, beer_shortdesc=?, created_at=? WHERE beer_id=?")
	if err != nil {
		return err
	}

	stmt.Exec(b.Name, b.Brewery, b.Abv, b.ShortDesc, b.Created, b.ID)

	return nil
}

func (s Storage) BeerDelete(id int) error {
	delBeer, err := s.db.Prepare("DELETE FROM beer WHERE beer_id=?")
	if err != nil {
		return err
	}

	delBeer.Exec(id)

	return nil
}
