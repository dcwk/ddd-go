package aggregate

import (
	"errors"

	"github.com/dcwk/ddd-go/entity"
	"github.com/dcwk/ddd-go/valueobject"
	"github.com/google/uuid"
)

var ErrInvalidPerson = errors.New("a customer has to have an valid person")

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []*valueobject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
			person:       person,
			products:     make([]*entity.Item, 0),
			transactions: make([]*valueobject.Transaction, 0),
		},
		nil
}

func (c *Customer) GetId() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetId(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.ID = id
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.Name = name
}
