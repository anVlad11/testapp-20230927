package services

import "github.com/labstack/echo/v4"

type HandlerService interface {
	CreateOrder(c echo.Context) error
}
