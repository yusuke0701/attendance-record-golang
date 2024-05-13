package router

import "github.com/labstack/echo/v4"

type echoEngine struct {
	router *echo.Echo
	port   Port
}

func newEchoServer(port Port) *echoEngine {
	return &echoEngine{
		router: echo.New(),
		port:   port,
	}
}

func (e echoEngine) Listen() {
	// TODO: Implement this method
}
