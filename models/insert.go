package models

import "testando-go/db"

// função que caracteriza a inserção de um to-do no banco de dados
func insert(todo Todo) (id int64, err error) {

	// abre a conexão com o banco de dados
	conn, err := db.OpenConnection()

	// se houver erro, será tratado pelo handler
	if err != nil {
		// haverá um devido handler pra essa situação futuramente
		return
	}
	defer conn.Close()

	// executa a query de insert
	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	// QueryRow retorna uma única linha de uma query
	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
