package graph

import "github.com/josesmar/20-clean-arch/internal/usecase"

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	GetOrderUseCase    usecase.GetOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}
