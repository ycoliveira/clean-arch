package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type ListOrdersOutputDTO struct {
	Items []OrderOutputDTO `json:"items"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{OrderRepository: OrderRepository}
}

func (c *ListOrdersUseCase) Execute() (*ListOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.ListAll()
	if err != nil {
		return nil, err
	}

	var items []OrderOutputDTO
	for _, order := range orders {
		item := OrderOutputDTO{ID: order.ID, Price: order.Price, Tax: order.Tax, FinalPrice: order.FinalPrice}
		items = append(items, item)
	}

	return &ListOrdersOutputDTO{Items: items}, nil
}
