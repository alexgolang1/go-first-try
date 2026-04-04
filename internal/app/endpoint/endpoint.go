package endpoint

import (
	"back-api/internal/app/types"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	GetUserByID(ctx context.Context, id int) (*types.Model, error)
	CreateUser(name, surname string) error
	DeleteUser(id int) error
	UpdateUser(ctx context.Context, person types.Model) error
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
	var context context.Context
	id, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := end.repo.GetUserByID(context, id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
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
	var person types.Model
	var context context.Context

	if err := end.repo.UpdateUser(context, person); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, nil)
}
