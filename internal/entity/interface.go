package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	FindByID(id string) (*Order, error)
	FindAll() ([]Order, error)
}
