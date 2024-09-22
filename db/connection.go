package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"testando-go/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	conn, err := sql.Open("postgres", sc)
	if err != nil {
		// não é saudável tratar o erro dessa forma em produção
		panic(err)
	}
	err = conn.Ping()
	return conn, err
}
