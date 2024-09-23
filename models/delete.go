package models

import "testando-go/db"

// função delete que executa o comando SQL de deleção de um to-do
func Delete(id int64, todo Todo) (int64, error) {
	conn, err := db.OpenConnection()

	// retorno de erro caso haja problemas
	if err != nil {
		// SERÁ TRATADO PELO HANDLER
		return 0, err
	}

	// defer é usado para garantir que a conexão será fechada ao final da função
	defer conn.Close()

	// executa o comando SQL de deleção de um to-do
	res, err := conn.Exec(`DELETE FROM todos WHERE id = $1`, id)

	// retorno de erro caso haja problemas
	if err != nil {
		// haverá um devido handler pra esse problema futuramente
		return 0, err
	}

	// RowsAffected retorna o número de linhas afetadas para debbuging
	return res.RowsAffected()
}
