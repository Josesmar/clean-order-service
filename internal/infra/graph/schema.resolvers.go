package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.56

import (
	"context"
	"log"

	"github.com/josesmar/20-clean-arch/internal/infra/graph/model"
	"github.com/josesmar/20-clean-arch/internal/usecase"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.OrderInput) (*model.Order, error) {
	dto := usecase.OrderInputDTO{
		ID:    input.ID,
		Price: float64(input.Price),
		Tax:   float64(input.Tax),
	}
	output, err := r.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &model.Order{
		ID:         output.ID,
		Price:      float64(output.Price),
		Tax:        float64(output.Tax),
		FinalPrice: float64(output.FinalPrice),
	}, nil
}

// GetOrder is the resolver for the getOrder field.
func (r *queryResolver) GetOrder(ctx context.Context, id string) (*model.Order, error) {
	order, err := r.GetOrderUseCase.Execute(id)
	if err != nil {
		log.Println("Erro ao buscar pedido:", err) // Adicionando log para captura de erro
		return nil, err
	}
	log.Println("Pedido encontrado:", order) // Log de sucesso
	return &model.Order{
		ID:         order.ID,
		Price:      float64(order.Price),
		Tax:        float64(order.Tax),
		FinalPrice: float64(order.FinalPrice),
	}, nil
}

// ListOrders is the resolver for the listOrders field.
func (r *queryResolver) ListOrders(ctx context.Context) ([]*model.Order, error) {
	orders, err := r.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var result []*model.Order
	for _, order := range orders {
		result = append(result, &model.Order{
			ID:         order.ID,
			Price:      float64(order.Price),
			Tax:        float64(order.Tax),
			FinalPrice: float64(order.FinalPrice),
		})
	}
	return result, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }