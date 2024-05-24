package usecase

import (
	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type ListOrderOutputDTO struct {
	Orders []OrderOutputDTO
}

type GetOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *GetOrderUseCase {
	return &GetOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *GetOrderUseCase) Execute() (ListOrderOutputDTO, error) {
	orders, err := c.OrderRepository.GetAll()
	if err != nil {
		return ListOrderOutputDTO{}, err
	}

	var ordersDTO []OrderOutputDTO
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
		ordersDTO = append(ordersDTO, dto)
	}

	return ListOrderOutputDTO{Orders: ordersDTO}, nil
}