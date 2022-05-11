package service

import (
	"levanidus/todo-app"
	"levanidus/todo-app/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)

	// list doesnt exit or belong to user
	if err != nil {
		return 0, err
	}

	return s.repo.TodoItem.Create(listId, item)

}
