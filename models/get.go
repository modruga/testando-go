package models

import "testando-go/db"

// get é uma função que retorna um to-do específico do banco de dados
func get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()

	// retorno de erro caso haja problemas
	if err != nil {
		// haverá um devido handler pra essa situação futuramente
		return
	}
	defer conn.Close()

	// QueryRow retorna uma única linha de uma query
	row := conn.QueryRow(`SELECT id, title, description, done FROM todos WHERE id = $1`, id)

	// Scan copia os valores das colunas da linha para os valores passados como argumento
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
	return
}
