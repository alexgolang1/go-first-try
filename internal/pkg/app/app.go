package app

import (
	"back-api/internal/app/db"
	"back-api/internal/app/endpoint"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	repo *db.DataBase
	end  *endpoint.Endpoint
	echo *echo.Echo
}

func New(database *gorm.DB) *App {
	repo := db.NewDbase(database)
	end := endpoint.New(repo)
	echo := echo.New()

	return &App{
		end:  end,
		echo: echo,
	}
}

func (app *App) Run() {
	app.echo.GET("/get/:id", app.end.ID)
	app.echo.POST("/post", app.end.Create)
	app.echo.Logger.Fatal(app.echo.Start(":8080"))
}
