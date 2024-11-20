package usecase

import (
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
	// if g.OrderRepository == nil {
	// 	return nil, fmt.Errorf("OrderRepository is nil")
	// }

	order, err := g.OrderRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}
