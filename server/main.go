package main

import (
	"attendance-record/model"
	"attendance-record/router"

	"github.com/labstack/echo/v4"
)

func main() {
	sqlDB := model.DBConnection()
	defer sqlDB.Close()

	e := echo.New()
	router.SetRouter(e)
}
