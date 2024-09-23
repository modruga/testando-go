package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"testando-go/configs"
)

// função que retorna uma conexão com o banco de dados via ponteiro, ou erro caso encontre problemas
func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	// string de conexão
	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.Database)

	// abre a conexão com o banco de dados
	conn, err := sql.Open("postgres", sc)
	// retorno de erro caso haja problemas
	if err != nil {
		// haverá um devido handler pra esse problema futuramente
		panic(err)
	}
	err = conn.Ping()
	return conn, err
}
