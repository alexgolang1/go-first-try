package endpoint

import (
	"back-api/internal/app/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	GetUser(id int) (*types.Model, error)
	CreateUser(name, surname string) error
	DeleteUser(id int) error
	UpdateUser(id int, name string, surname string) error
}

type Endpoint struct {
	repo Repository
}

func New(r Repository) *Endpoint {
	return &Endpoint{
		repo: r,
	}
}

func (end *Endpoint) ID(ctx echo.Context) error {
	idstr := ctx.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := end.repo.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, result)
}

func (end *Endpoint) Create(ctx echo.Context) error {
	name := ctx.Param("Name")
	surname := ctx.Param("Surname")
	err := end.repo.CreateUser(name, surname)
	if err != nil {
		return ctx.JSON(500, err)
	}
	return ctx.JSON(http.StatusCreated, nil)
}

func (end *Endpoint) Delete(ctx echo.Context) error {
	idstr := ctx.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	if err := end.repo.DeleteUser(id); err != nil {
		ctx.JSON(http.StatusNotFound, err)
	}
	return ctx.JSON(http.StatusOK, nil)
}

func (end *Endpoint) Update(ctx echo.Context) error {
	idstr := ctx.Param("id")
	id, _ := strconv.Atoi(idstr)

	name := ctx.Param("Name")
	surname := ctx.Param("Surname")

	if err := end.repo.UpdateUser(id, name, surname); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, nil)
}
