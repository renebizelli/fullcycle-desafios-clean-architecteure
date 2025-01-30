package entity

import "github.com/devfullcycle/20-CleanArch/internal/dtos"

type OrderRepositoryInterface interface {
	Save(order *Order) error
	ListOrders() ([]dtos.OrderOutputDTO, error)
}
