package project

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"todoListApp/delivery/helper"
	"todoListApp/delivery/middlewares"
	"todoListApp/entities"
	"todoListApp/usecase/project"
)

type ProjectHandler struct {
	projectUseCase project.ProjectUseCaseInterface
}

func NewProjectHandler(projectUseCase project.ProjectUseCaseInterface) *ProjectHandler {
	return &ProjectHandler{
		projectUseCase: projectUseCase,
	}
}

func (ph *ProjectHandler) GetAllProjectByUserId() echo.HandlerFunc {
	return func(c echo.Context) error {

		var id = middlewares.ExtractToken(c)

		users, err := ph.projectUseCase.GetAllProjectByUserId(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get project by user id", users))
	}
}

func (ph *ProjectHandler) CreateProject() echo.HandlerFunc {
	return func(c echo.Context) error {
		var project entities.Project
		c.Bind(&project)
		id := middlewares.ExtractToken(c)
		project.TodoID = uint(id)
		err := ph.projectUseCase.CreateProject(project)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to create prject"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success Create project"))
	}
}

func (ph *ProjectHandler) DeleteProject() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		var userID = middlewares.ExtractToken(c)

		err := ph.projectUseCase.DeleteProject(id, userID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success delete project"))
	}
}

func (ph *ProjectHandler) UpdateProject() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id, _ = strconv.Atoi(c.Param("id"))
		var project entities.Project
		c.Bind(&project)

		err := ph.projectUseCase.UpdateProject(id, project)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update project", project))
	}
}
