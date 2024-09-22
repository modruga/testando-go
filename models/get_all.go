package models

import "testando-go/db"

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		// SERÁ TRATADO PELO HANDLER
		return
	}
	defer conn.Close()

	// SELECT * FROM todos
	rows, err := conn.Query(`SELECT * FROM todos`)

	// se houver erro, será tratado pelo handler
	if err != nil {
		// SERÁ TRATADO PELO HANDLER
		return
	}

	// se não houver erro, escaneia todos os retornos...
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}

		// ...e faz append ao slice
		todos = append(todos, todo)
	}

	return
}
