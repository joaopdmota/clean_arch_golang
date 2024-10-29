package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type GetOrdersOutputDTO struct {
	Orders []entity.Order `json:"data"`
}

type GetOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *GetOrdersUseCase) Execute() (GetOrdersOutputDTO, error) {
	result, err := c.OrderRepository.GetAll()

	if err != nil {
		return GetOrdersOutputDTO{
			Orders: []entity.Order{},
		}, err
	}

	return GetOrdersOutputDTO{
		Orders: result,
	}, nil
}
