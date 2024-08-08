package usecase

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UnitOfWork interface {
	Do(ctx context.Context, f func(tx *sqlx.Tx) error) error
}

type unitOfWork struct {
	db *sqlx.DB
}

func NewUnitOfWork(db *sqlx.DB) UnitOfWork {
	return &unitOfWork{
		db: db,
	}
}

func (u *unitOfWork) Do(ctx context.Context, f func(tx *sqlx.Tx) error) error {
	tx, err := u.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
        if v := recover(); v != nil {
			if rerr := tx.Rollback(); rerr != nil {
				err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
			}
            panic(v)
        }
    }()
    if err := f(tx); err != nil {
        if rerr := tx.Rollback(); rerr != nil {
            err = fmt.Errorf("%w: rolling back transaction: %v", err, rerr)
        }
        return err
    }
    if err := tx.Commit(); err != nil {
        return fmt.Errorf("committing transaction: %w", err)
    }

	return nil
}

