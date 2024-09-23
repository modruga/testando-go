package models

import "testando-go/db"

// GetAll é uma função que retorna todos os to-dos do banco de dados
func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()

	// caso haja erro, será tratado pelo handler
	if err != nil {
		// haverá um devido handler pra essa situação futuramente
		return
	}
	defer conn.Close()

	// SELECT * FROM todos
	rows, err := conn.Query(`SELECT * FROM todos`)

	// se houver erro, será tratado pelo handler
	if err != nil {
		// haverá um devido handler pra essa situação futuramente
		return
	}

	// se não houver erro, escaneia todos os dados to-do do banco de dados
	for rows.Next() {
		var todo Todo

		// Scan copia os valores das colunas da linha para os valores passados como argumento
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

		// se houver erro, continua o loop
		if err != nil {
			continue
		}

		// ...e faz append ao slice
		todos = append(todos, todo)
	}

	// retorna o slice de to-dos
	return
}
