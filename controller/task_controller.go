package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bookkyjung1221/echo-clean-architecture/model"
	"github.com/bookkyjung1221/echo-clean-architecture/usecase"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

func (tc *taskController) GetAllTasks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	tasksRes, err := tc.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res := &model.TaskResponseDetail{
		TaskResponse: tasksRes,
		Code:         http.StatusOK,
		Message:      "Get All Tasks Success",
	}

	return c.JSON(http.StatusOK, res)
}

func (tc *taskController) GetTaskById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)
	taskRes, err := tc.tu.GetTaskById(uint(userId.(float64)), uint(taskId))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.TaskResponseError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	res := &model.TaskResponseOneDetail{
		TaskResponse: taskRes,
		Code:         http.StatusOK,
		Message:      "Get Task By ID Success",
	}

	return c.JSON(http.StatusOK, res)
}

func (tc *taskController) CreateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	task := model.Task{}

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, &model.TaskResponseError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	task.UserId = uint(userId.(float64))
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.TaskResponseError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	res := &model.TaskResponseOneDetail{
		TaskResponse: taskRes,
		Code:         http.StatusCreated,
		Message:      "Create Task Success",
	}

	return c.JSON(http.StatusCreated, res)
}

func (tc *taskController) UpdateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	task := model.Task{}

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, &model.TaskResponseError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	taskRes, err := tc.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.TaskResponseError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	res := &model.TaskResponseOneDetail{
		TaskResponse: taskRes,
		Code:         http.StatusOK,
		Message:      "Update Task Success",
	}

	return c.JSON(http.StatusOK, res)
}

func (tc *taskController) DeleteTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	err := tc.tu.DeleteTask(uint(userId.(float64)), uint(taskId))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &model.TaskResponseError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	message := fmt.Sprintf("Delete taskId %s Success", id)

	res := &model.TaskResponseDelete{
		Code:    http.StatusOK,
		Message: message,
	}

	return c.JSON(http.StatusOK, res)
}
