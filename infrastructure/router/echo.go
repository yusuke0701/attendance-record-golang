package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"attendance-record/adapter/logger"
	"attendance-record/model"
)

type echoEngine struct {
	router *echo.Echo
	log    logger.Logger
	port   Port
}

func newEchoServer(log logger.Logger, port Port) *echoEngine {
	return &echoEngine{
		router: echo.New(),
		log:    log,
		port:   port,
	}
}

func (e echoEngine) Listen() {
	// e.router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}\n",
	// 	Output: os.Stdout,
	// }))
	e.router.Use(middleware.Recover())
	e.router.Use(middleware.CORS())

	e.setAppHandlers(e.router)

	if err := e.router.Start(":8000"); err != nil {
		e.log.Fatalln("failed to start server", err)
	}
}

func (e echoEngine) setAppHandlers(router *echo.Echo) {
	// TODO: ラップする
	router.GET("/api/tasks", GetTasksHandler)
	router.POST("/api/tasks", AddTasksHandler)
}

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
