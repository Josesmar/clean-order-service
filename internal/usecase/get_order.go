package usecase

import (
	"fmt"

	"github.com/josesmar/20-clean-arch/internal/entity"
)

type GetOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrderUseCase(orderRepository entity.OrderRepositoryInterface) *GetOrderUseCase {
	return &GetOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (g *GetOrderUseCase) Execute(id string) (*entity.Order, error) {

	order, err := g.OrderRepository.FindByID(id)
	if order == nil {
		return nil, fmt.Errorf("order with ID %s not found", id)
	}

	if err != nil {
		return nil, err
	}

	return order, nil
}
