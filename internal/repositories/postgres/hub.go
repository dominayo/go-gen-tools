package postgres

import (
	"context"

	"github.com/duyledat197/go-gen-tools/internal/models"
	"github.com/duyledat197/go-gen-tools/internal/repositories"
)

type hubRepository struct {
	db models.DBTX
}

func NewHubRepository(db models.DBTX) repositories.HubRepository {
	return &hubRepository{
		db: db,
	}
}

func (r *hubRepository) Create(ctx context.Context, hub *models.Hub, opts ...repositories.Options) error {
	// ToDo remove querier to arguments and make it isolative. Because we should'nt using a common model.querier
	q := models.New(r.db)
	return nil
}

func (r *hubRepository) Update(ctx context.Context, id string, hub *models.Hub, opts ...repositories.Options) error {
	q := models.New(r.db)
	return nil
}

func (r *hubRepository) Delete(ctx context.Context, id string, opts ...repositories.Options) error {
	q := models.New(r.db)
	return nil
}

func (r *hubRepository) GetList(ctx context.Context, offset, limit int, opts ...repositories.Options) ([]*models.Hub, error) {
	q := models.New(r.db)
	return nil, nil
}

func (r *hubRepository) GetByID(ctx context.Context, id string, opts ...repositories.Options) (*models.Hub, error) {
	q := models.New(r.db)
	return nil, nil
}
