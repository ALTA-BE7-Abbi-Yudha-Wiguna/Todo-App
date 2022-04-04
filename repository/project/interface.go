package project

import "todoListApp/entities"

type ProjectRepositoryInterface interface {
	GetAllProjectByUserId(UserID int) ([]entities.Project, error)
	CreateProject(project entities.Project) error
	DeleteProject(id, userID int) error
	UpdateProject(id int, project entities.Project) error
}
