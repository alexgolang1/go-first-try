package endpoint

import (
	"back-api/internal/app/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Endpoint struct {
	srv *service.Service
}

func New(srv *service.Service) *Endpoint {
	return &Endpoint{
		srv: srv,
	}
}

func (end *Endpoint) GetIDEnd(ctx echo.Context) error {
	idstr := ctx.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	result, err := end.srv.GetIDSrv(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, result)
}
