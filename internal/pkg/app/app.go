package app

import (
	"echo/middleware/internal/app/mv"
	endpoint "echo/middleware/internal/app/pkg"
	"echo/middleware/internal/app/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(mv.RoleCheck)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a *App) Run() error {
	fmt.Print("Server running")

	err := a.echo.Start(":4000")

	if err != nil {
		log.Fatal(err)
	}

	return nil
}
