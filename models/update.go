package models

import "testando-go/db"

// update é uma função que ATUALIZA um to-do no banco de dados
func update(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()

	// caso encontre erro, será tratado pelo handler
	if err != nil {
		// haverá um devido handler pra essa situação futuramente
		return 0, err
	}
	defer conn.Close()

	// executa o comando SQL de atualização de um to-do
	// Exec executa uma query que não retorna linhas
	res, err := conn.Exec(`UPDATE todos SET title = $2, description = $3, done = $3 WHERE id = $1`,
		id, todo.Title, todo.Description, todo.Done)

	// caso encontre erro, será tratado pelo handler
	if err != nil {
		// haverá um devido handler pra essa situação futuramente
		return 0, err
	}

	// RowsAffected retorna o número de linhas afetadas para debugging
	return res.RowsAffected()
}
