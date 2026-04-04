package auth

import (
	"back-api/internal/app/repository"
	"back-api/internal/app/types"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Login(ctx echo.Context) error {
	req := new(types.Model)
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	user, err := h.repo.GetUserByEmail(ctx.Request().Context(), req.Email)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, err)
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": token,
		"user":  user,
	})
}
