package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	status struct {
		db *sqlx.DB
	}
)

var _ repository.Status = (*status)(nil)

// Create accout repository
func NewStatus(db *sqlx.DB) *status {
	return &status{db: db}
}

func (s *status) Create(ctx context.Context, tx *sqlx.Tx, st *object.Status) error {
	_, err := tx.Exec("insert into status (account_id, content, create_at) values (?, ?, ?)", st.AccountID, st.Content, st.CreateAt)
	if err != nil {
		return fmt.Errorf("failed to insert status: %w", err)
	}

	return nil
}

func (s *status) FindByID(ctx context.Context, id string) (*object.Status, error) {
	entity := new(object.Status)
	err := s.db.QueryRowxContext(ctx, "select * from status where id = ?", id).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no status found with the given id: %w", err)
		}

		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}
	return entity, nil
}

func (s *status) FindPublicTimeline(ctx context.Context, limit int) ([]*object.Status, error) {
	var entities []*object.Status

	err := s.db.SelectContext(ctx, &entities, "select * from status order by id desc limit ?", limit)
	if err != nil {
		return nil, fmt.Errorf("failed to find status from db: %w", err)
	}
	return entities, nil
}