package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/matheusF23/pecunia-go/internal/main/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	err = conn.Ping()
	return conn, err
}
