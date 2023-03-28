package app

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/timvkim/middleware/internal/app/endpoint"
	"github.com/timvkim/middleware/internal/app/mw"
	"github.com/timvkim/middleware/internal/app/service"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	// app instanse
	a := &App{}

	// service instance
	a.s = service.New()

	// endpoint instance
	a.e = endpoint.New(a.s)

	// server instance
	a.echo = echo.New()

	// middleware for checking a role
	a.echo.Use(mw.RoleCheck)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {

	fmt.Println("server running")

	// start HTTP server and error handling
	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
