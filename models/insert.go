package models

import "testando-go/db"

func insert(todo Todo) (id int64, err error) {

	// começa abrindo a conexão
	conn, err := db.OpenConnection()

	// se houver erro, será tratado pelo handler
	if err != nil {
		// SERÁ TRATADO PELO HANDLER
		return
	}
	defer conn.Close()

	// executa a query de insert
	sql := `INSERT INTO todos (title, description, done) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, todo.Title, todo.Description, todo.Done).Scan(&id)

	return
}
