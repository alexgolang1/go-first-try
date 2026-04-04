package app

import (
	"back-api/internal/app/auth"
	"back-api/internal/app/endpoint"
	"back-api/internal/app/mw"
	"back-api/internal/app/repository"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	repo *repository.Repository
	end  *endpoint.Endpoint
	echo *echo.Echo
	auth *auth.Handler
}

func New(database *gorm.DB, rdb *redis.Client) *App {
	repo := repository.NewRepository(database, rdb)
	auth := auth.NewHandler(repo)
	end := endpoint.New(repo)
	echo := echo.New()

	return &App{
		repo: repo,
		auth: auth,
		end:  end,
		echo: echo,
	}
}

func (app *App) Run() {
	app.echo.GET("/get/:id", app.end.ID, mw.MiddleWare)
	app.echo.POST("/post", app.end.Create)
	app.echo.POST("/login", app.auth.Login)
	app.echo.DELETE("/delete/:id", app.end.Delete)
	app.echo.PUT("/update/:id", app.end.Create)
	app.echo.Logger.Fatal(app.echo.Start(":8080"))
}
