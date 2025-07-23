package tools

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type postgresDB struct {
	DB *sql.DB
}

func (pg *postgresDB) SetupDatabase() error {
	connStr := "postgres://postgres@localhost/goapi?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	pg.DB = db
	fmt.Println(" Connected to PostgreSQL")
	return nil
}

func (pg *postgresDB) GetUserLoginDetails(username string) *loginDetails {
	row := pg.DB.QueryRow("SELECT username, auth_token FROM login_details WHERE username = $1", username)

	var user loginDetails
	if err := row.Scan(&user.Username, &user.AuthToken); err != nil {
		log.Println("User not found:", err)
		return nil
	}
	return &user
}

func (pg *postgresDB) GetUserCoins(username string) *CoinDetails {
	row := pg.DB.QueryRow("SELECT username, coins FROM coin_details WHERE username = $1", username)

	var coin CoinDetails
	if err := row.Scan(&coin.Username, &coin.Coins); err != nil {
		log.Println("Coin details not found:", err)
		return nil
	}
	return &coin
}
