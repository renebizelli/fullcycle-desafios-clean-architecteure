package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/dtos"
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]dtos.OrderOutputDTO, error) {

	orders, err := c.OrderRepository.ListOrders()
	if err != nil {
		return nil, err
	}

	return orders, nil
}
