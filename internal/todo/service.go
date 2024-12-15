package todo

func GetTodos() ([]Todo, error) {
	return fetchAllTodos()
}

func GetTodoByID(id int) (*Todo, error) {
	return getOneTodo(id)
}

func CreateTodo(title string) error {
	return insertTodo(title)
}

func UpdateTodo(id int, title string) error {
	return updateTodo(id, title)
}

func CompleteTodoByID(id int) error {
	return markTodoCompleted(id)
}

func DeleteTodoByID(id int) error {
	return deleteTodoByID(id)
}
