package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

const (
	host   = "localhost"
	port   = 5432
	dbName = "todo"
)

func GetPostgresDatabase() (db *sql.DB, err error) {
	user := viper.Get("postgres.user")
	password := viper.Get("postgres.key")
	connectionInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err = sql.Open("postgres", connectionInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping() // force a call to open a connection
	if err != nil {
		panic(err)
	}

	return db, nil
}
