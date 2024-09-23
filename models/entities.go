package models

// definição da struct to-do, que será usada para representar uma tabela no banco de dados
type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
