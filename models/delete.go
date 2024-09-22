package models

import "testando-go/db"

func Delete(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		// SERÁ TRATADO PELO HANDLER
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM todos WHERE id = $1`, id)

	if err != nil {
		// SERÁ TRATADO PELO HANDLER
		return 0, err
	}

	// RowsAffected retorna o número de linhas afetadas
	return res.RowsAffected()
}
