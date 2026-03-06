package endpoint

import (
	"back-api/internal/app/types"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	GetID(id int) (*types.Model, error)
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

	result, err := end.repo.GetID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, result)
}
