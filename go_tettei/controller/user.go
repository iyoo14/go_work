package controller

import (
	"fmt"
	"go_work/go_tettei/usecase"
	"net/http"
)
import "github.com/labstack/echo"

type UserController interface {
	Get(ctx echo.Context) error
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}

type userController struct {
	u usecase.UserUsecase
}

func NewUserController(u usecase.UserUsecase) UserController {
	return &userController{u}
}

func (c *userController) Get(ctx echo.Context) error {
	id := ctx.Param("id")
	u, err := c.u.GetByID(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, u)
}

func (c *userController) Create(ctx echo.Context) error {
	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	u := toModel(req)
	fmt.Println("creating...", req)
	fmt.Println(u.Name, u.Age, u.Email, u.CreatedAt, u.UpdatedAt)
	id, err := c.u.Create(ctx.Request().Context(), u)
	if err != nil {
		fmt.Println("creating error...", err)
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, id)
}

func (c *userController) Update(ctx echo.Context) error {
	var req UserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := ctx.Validate(req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	u := toModel(req)
	c.u.Update(ctx.Request().Context(), u)
	return nil
}

func (c *userController) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	err := c.u.Delete(ctx.Request().Context(), id)
	if err != nil {
		fmt.Println("delete error...", err)
		return ctx.JSON(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, id)
}
