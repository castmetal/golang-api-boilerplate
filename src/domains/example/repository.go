package example

import (
	"context"

	"github.com/castmetal/golang-api-boilerplate/src/domains/common"
	"github.com/google/uuid"
)

type IExampleRepository interface {
	common.IAggregateRoot
	Create(ctx context.Context, example *Example) error
	FindOneById(ctx context.Context, id uuid.UUID) (*Example, error)
	ListAll(ctx context.Context, limit int, offset int) ([]*Example, error)
	FindOneByName(ctx context.Context, name string) (*Example, error)
}
