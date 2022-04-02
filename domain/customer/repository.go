package customer

import (
	"errors"

	"github.com/dcwk/ddd-go/aggregate"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	ErrUpdatedCustomer     = errors.New("failed to update the customer in the repository")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
