package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	// TODO: Add Other APIs
	Create(ctx context.Context, tx *sqlx.Tx, st *object.Status) error
	FindByID(ctx context.Context, id string) (*object.Status, error)
	FindPublicTimeline(ctx context.Context, limit int) ([]*object.Status, error)
}
