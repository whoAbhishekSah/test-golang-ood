package controller2

import "todo-golang/model"

type AnotherTypeOfController struct{
}

func (t * AnotherTypeOfController) Create(string, string) (*model.ToDo, error){
	return &model.ToDo{}, nil
}

func (t *AnotherTypeOfController) Creadsfsfsdfsdfsdfasdfsde(string, string) (*model.ToDo, error){
	return &model.ToDo{}, nil
}

func (t *AnotherTypeOfController) SearchById(int64) (*model.ToDo, error){
	return &model.ToDo{}, nil
}

func (t *AnotherTypeOfController) SearchByTitle(string)(*[]model.ToDo, error){
	return nil, nil
}

func NewAnotherTypeOfController() *AnotherTypeOfController {
	return &AnotherTypeOfController{}
}
