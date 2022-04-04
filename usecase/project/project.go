package project

import (
	"todoListApp/entities"
	"todoListApp/repository/project"
)

type ProjectUseCase struct {
	ProjectRepository project.ProjectRepositoryInterface
}

func NewProjectUseCase(projectRepo project.ProjectRepositoryInterface) ProjectUseCaseInterface {
	return &ProjectUseCase{
		ProjectRepository: projectRepo,
	}

}

func (puc *ProjectUseCase) GetAllProjectByUserId(UserID int) ([]entities.Project, error) {
	projects, err := puc.ProjectRepository.GetAllProjectByUserId(UserID)
	return projects, err
}

func (puc *ProjectUseCase) CreateProject(project entities.Project) error {
	err := puc.ProjectRepository.CreateProject(project)
	return err
}

func (puc *ProjectUseCase) DeleteProject(id, userID int) error {
	err := puc.ProjectRepository.DeleteProject(id, userID)
	return err
}

func (puc *ProjectUseCase) UpdateProject(id int, project entities.Project) error {
	err := puc.ProjectRepository.UpdateProject(id, project)
	return err
}
