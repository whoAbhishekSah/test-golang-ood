package controller

import "todo-golang/model"

type Controller interface {
    Create(string, string) (*model.ToDo, error)
    SearchById(int64) (*model.ToDo, error)
    SearchByTitle(string)(*[]model.ToDo, error)
}


type TodoController struct {
	model model.TodoModel
}

func (t TodoController) Create(string, string) (*model.ToDo, error){
	return &model.ToDo{}, nil
}

func (t TodoController) Creadsfsfsdfsdfsdfasdfsde(string, string) (*model.ToDo, error){
	return &model.ToDo{}, nil
}

func (t TodoController) SearchById(int64) (*model.ToDo, error){
	return &model.ToDo{}, nil
}

func (t TodoController) SearchByTitle(string)(*[]model.ToDo, error){
	return nil, nil
}

func NewTodoController() (Controller) {
	return TodoController{ 
		model: model.NewDBBASEDModel(),
		}
}
