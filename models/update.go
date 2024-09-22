package models

import "testando-go/db"

func update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		// SERÁ TRATADO PELO HANDLER
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE todos SET title = $2, description = $3, done = $3 WHERE id = $1`,
		id, todo.Title, todo.Description, todo.Done)
	if err != nil {
		// SERÁ TRATADO PELO HANDLER
		return 0, err
	}

	// RowsAffected retorna o número de linhas afetadas
	return res.RowsAffected()
}
