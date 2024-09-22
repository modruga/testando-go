package models

import "testando-go/db"

func get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		// SERÁ TRATADO PELO HANDLER
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT id, title, description, done FROM todos WHERE id = $1`, id)
	err = row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)
	return
}
