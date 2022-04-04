package todo

import (
	"fmt"
	"gorm.io/gorm"
	"todoListApp/entities"
)

type TodoRepository struct {
	database *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{
		database: db,
	}
}

func (tr *TodoRepository) GetAllTodos() ([]entities.Todo, error) {
	var todos []entities.Todo
	tx := tr.database.Find(&todos)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return todos, nil

}

func (tr *TodoRepository) GetTodoById(id int) (entities.Todo, error) {
	var todos entities.Todo

	tx := tr.database.Where("id = ?", id).First(&todos)
	if tx.Error != nil {
		return todos, tx.Error
	}
	return todos, nil

}

func (tr *TodoRepository) CreateTodo(todo entities.Todo) error {

	tx := tr.database.Save(&todo)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (tr *TodoRepository) DeleteTodo(id, userID int) error {
	var todos entities.Todo
	err := tr.database.Where("id =? and user_id = ?", id, userID).First(&todos).Error
	if err != nil {
		return err
	}
	tr.database.Delete(&todos)

	return nil

}

func (tr *TodoRepository) UpdateTodo(id int, todo entities.Todo) error {

	tx := tr.database.Where("id = ?", id).Updates(&todo)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {

		return fmt.Errorf("%s", "error")
	}
	return nil

}

func (tr *TodoRepository) CompleteTodo(todo entities.Todo) error {

	tx := tr.database.Where("id = ?", todo.ID).Updates(&todo)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (tr *TodoRepository) ReOpen(todo entities.Todo) error {

	tx := tr.database.Where("id = ?", todo.ID).Updates(&todo)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (tr *TodoRepository) GetAllTodosByIdUser(UserID int) ([]entities.Todo, error) {
	var todos []entities.Todo

	tx := tr.database.Where("user_id = ?", UserID).Find(&todos)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return todos, nil

}
