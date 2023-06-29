package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/castmetal/golang-api-boilerplate/src/domains/common"
	"github.com/castmetal/golang-api-boilerplate/src/domains/example"
	"github.com/castmetal/golang-api-boilerplate/src/infra/storage/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type exampleRepository struct {
	db        *pgxpool.Pool
	pgQueries *postgres.Queries
}

func NewExampleRepository(db *pgxpool.Pool) example.IExampleRepository {
	q := postgres.New(db)

	return &exampleRepository{db: db, pgQueries: q}
}

func (r *exampleRepository) Create(ctx context.Context, example *example.Example) error {
	return r.pgQueries.CreateExample(ctx, postgres.CreateExampleParams{
		ID:        example.ID,
		Name:      example.Name,
		CreatedAt: example.CreatedAt.Value,
		UpdatedAt: example.UpdatedAt.Value,
		DeletedAt: sql.NullTime{Time: time.Time{}, Valid: false},
	})
}

func (r *exampleRepository) FindOneById(ctx context.Context, id uuid.UUID) (*example.Example, error) {
	exmpl, err := r.pgQueries.GetExampleByID(ctx, id)
	if err != nil || exmpl.ID.String() == "" {
		return nil, common.NotFoundError("The example id " + id.String())
	}

	mapper := r.ExampleMapper(exmpl)

	return mapper, nil
}

func (r *exampleRepository) FindOneByName(ctx context.Context, name string) (*example.Example, error) {
	exmpl, err := r.pgQueries.GetExampleByName(ctx, name)
	if err != nil || exmpl.ID.String() == "" {
		return nil, common.NotFoundError("The example: name " + name)
	}

	mapper := r.ExampleMapper(exmpl)

	return mapper, nil
}

func (r *exampleRepository) ListAll(ctx context.Context, limit int, offset int) ([]*example.Example, error) {
	exmpls, err := r.pgQueries.ListAllExamples(ctx, postgres.ListAllExamplesParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, err
	}

	mapper := r.ExampleMapperList(exmpls)

	return mapper, nil
}

func (r *exampleRepository) ExampleMapper(data postgres.Example) *example.Example {
	var exmpl *example.Example

	exmpl = &example.Example{
		ID:        data.ID,
		Name:      data.Name,
		CreatedAt: common.JsonTime{Value: data.CreatedAt},
		UpdatedAt: common.JsonTime{Value: data.UpdatedAt},
	}

	if data.DeletedAt.Valid && data.DeletedAt.Time.String() != "" {
		exmpl.DeletedAt = &common.JsonNullTime{Value: data.DeletedAt}
	}

	return exmpl
}

func (r *exampleRepository) ExampleMapperList(data []postgres.Example) []*example.Example {
	var exmpls []*example.Example

	for _, value := range data {
		exmpl := r.ExampleMapper(value)
		exmpls = append(exmpls, exmpl)
	}

	return exmpls
}
