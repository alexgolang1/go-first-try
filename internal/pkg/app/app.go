package app

import (
	"back-api/internal/app/db"
	"back-api/internal/app/endpoint"
	"back-api/internal/app/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	repo *db.Repository
	srv  *service.Service
	end  *endpoint.Endpoint
	echo *echo.Echo
}

func New(database *gorm.DB) *App {
	repos := db.NewRepo(database)
	serve := service.New(repos)
	endpoin := endpoint.New(serve)
	ech := echo.New()

	return &App{
		srv:  serve,
		end:  endpoin,
		echo: ech,
	}
}

func (app *App) Run() {
	app.echo.GET("/anyurl", app.end.GetIDEnd)
	app.echo.Logger.Fatal(app.echo.Start(":8080"))
}
