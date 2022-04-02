package memory

import (
	"sync"

	"github.com/dcwk/ddd-go/aggregate"
	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{}
}
