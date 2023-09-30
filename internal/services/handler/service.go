package handler

import (
	"github.com/anvlad11/testapp-20230927/internal/interfaces/services"
	"github.com/anvlad11/testapp-20230927/pkg/errors"
	contract "github.com/anvlad11/testapp-20230927/pkg/testapp"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service struct {
	orderService services.OrderService
}

func NewService(orderService services.OrderService) *Service {
	return &Service{
		orderService: orderService,
	}
}

func (s *Service) CreateOrder(c echo.Context) error {
	var request *contract.V1CreateOrderRequestBody

	err := c.Bind(&request)
	if err != nil {
		response := &contract.V1CreateOrderErrorResponseBody{Error: errors.RequestBodyIsMalformed.Error()}

		return c.JSON(http.StatusBadRequest, response)
	}

	if request.ItemsCount <= 0 {
		response := &contract.V1CreateOrderErrorResponseBody{Error: errors.ItemCountIsEqualOrLessThanZero.Error()}

		return c.JSON(http.StatusBadRequest, response)
	}

	var packs map[int]int
	packs, err = s.orderService.CreateOrder(request.ItemsCount)

	if err != nil {
		response := &contract.V1CreateOrderErrorResponseBody{Error: errors.InternalError.Error()}

		return c.JSON(http.StatusInternalServerError, response)
	}

	response := &contract.V1CreateOrderSuccessfulResponseBody{Packs: make([]contract.V1Pack, 0, len(packs))}
	for size, count := range packs {
		response.Packs = append(response.Packs, contract.V1Pack{
			Count: count,
			Size:  size,
		})
	}

	return c.JSON(http.StatusOK, response)
}
