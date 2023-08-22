package handler

import "fmt"

type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Completed   bool   `json:"completed" default:"false"`
}

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Completed   bool   `json:"completed" default:"false"`
}

type UpdateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s is required and must be %s", name, typ)
}

func (r *CreateTodoRequest) Validate() error {
	if r.Title == "" && r.Description == "" {
		return fmt.Errorf("request body cannot be empty")
	}
	if r.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if r.Description == "" {
		return errParamIsRequired("description", "string")
	}
	return nil
}

func (r *UpdateTodoRequest) Validate() error {
	if r.Title != "" || r.Description != "" {
		return nil
	}
	return fmt.Errorf("at least one field must be filled")
}

func (r *UpdateTodoRequest) Merge(todo *Todo) {
	if r.Title != "" {
		todo.Title = r.Title
	}
	if r.Description != "" {
		todo.Description = r.Description
	}
	if r.Completed {
		todo.Completed = r.Completed
	}
}

func (r *UpdateTodoRequest) ToTodo() *Todo {
	return &Todo{
		Title:       r.Title,
		Description: r.Description,
		Completed:   r.Completed,
	}
}
