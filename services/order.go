package services

import (
	"github.com/dcwk/ddd-go/domain/customer"
	"github.com/dcwk/ddd-go/domain/customer/memory"
	"github.com/gofrs/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) error {
	_, err := o.customers.Get(customerID)
	if err != nil {
		return err
	}

	return nil
}
