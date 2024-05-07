package router

import (
	"net/http"

	"attendance-record/model"

	"github.com/labstack/echo/v4"
)

func GetTasksHandler(c echo.Context) error {
	tasks, err := model.GetTasks()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.JSON(http.StatusOK, tasks)
}

type ReqTask struct {
	Name string `json:"name"`
}

func AddTasksHandler(c echo.Context) error {
	req := new(ReqTask)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}

	task, err := model.AddTask(req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Bad Request")
	}
	return c.JSON(http.StatusOK, task)
}
