package service

import (
	"context"

	"github.com/josesmar/20-clean-arch/internal/infra/grpc/pb"
	"github.com/josesmar/20-clean-arch/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	GetOrderUseCase    usecase.GetOrderUseCase
	ListOrderUseCase   usecase.ListOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase,
	getOrderUseCase usecase.GetOrderUseCase, listOrderUseCase usecase.ListOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		GetOrderUseCase:    getOrderUseCase,
		ListOrderUseCase:   listOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) GetOrder(ctx context.Context, in *pb.GetOrderRequest) (*pb.OrderResponse, error) {
	id := in.GetId()
	output, err := s.GetOrderUseCase.Execute(id)
	if err != nil {
		return nil, err
	}
	return &pb.OrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.OrdersListResponse, error) {
	orders, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}
	var result []*pb.OrderResponse
	for _, order := range orders {
		result = append(result, &pb.OrderResponse{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}
	return &pb.OrdersListResponse{Orders: result}, nil
}
