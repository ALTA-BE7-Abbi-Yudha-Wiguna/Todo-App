package project

import (
	"fmt"
	"gorm.io/gorm"
	"todoListApp/entities"
)

type ProjectRepository struct {
	database *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		database: db,
	}
}

func (pr *ProjectRepository) GetAllProjectByUserId(UserID int) ([]entities.Project, error) {
	var projects []entities.Project
	tx := pr.database.Where("user_id = ?", UserID).Find(&projects)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return projects, nil

}

func (pr *ProjectRepository) CreateProject(project entities.Project) error {

	tx := pr.database.Save(&project)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

func (pr *ProjectRepository) DeleteProject(id, userID int) error {
	var project entities.Project
	err := pr.database.Where("id =? and user_id = ?", id, userID).First(&project).Error
	if err != nil {
		return err
	}
	pr.database.Delete(&project)

	return nil

}

func (pr *ProjectRepository) UpdateProject(id int, project entities.Project) error {

	tx := pr.database.Where("id = ?", id).Updates(&project)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {

		return fmt.Errorf("%s", "error")
	}
	return nil

}
