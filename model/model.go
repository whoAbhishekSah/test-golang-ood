package model

type TodoModel interface{
	CreateTodo(string, string) (*ToDo, error)
}

type ToDo struct{
	id int64
	title string
}


func NewDBBASEDModel() ToDo{
	return ToDo{}
}
func (t ToDo) CreateTodo(string, string) (*ToDo, error) {
	return nil, nil
}
