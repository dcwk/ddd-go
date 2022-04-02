package aggregate

import "github.com/dcwk/ddd-go/entity"

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []*valueobject.Transaction
}
