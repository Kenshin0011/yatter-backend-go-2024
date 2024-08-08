package usecase

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type Status interface {
	Create(ctx context.Context, account_id int, content string) (*CreateStatusDTO, error)
	FindByID(ctx context.Context, id string) (*GetStatusDTO, error)
	FindPublicTimeline(ctx context.Context, limit int) ([]*object.Status, error)
}

type status struct {
	db          *sqlx.DB
	statusRepo 	repository.Status
	unitOfWork  UnitOfWork
}

type CreateStatusDTO struct {
	Status *object.Status
}

type GetStatusDTO struct {
	Status *object.Status
}

var _ Status = (*status)(nil)

func NewStatus(db *sqlx.DB, statusRepo repository.Status, unitOfWork UnitOfWork) *status {
	return &status{
		db:          db,
		statusRepo:  statusRepo,
		unitOfWork:  unitOfWork,
	}
}

func (s *status) Create(ctx context.Context, account_id int,content string) (*CreateStatusDTO, error) {
	st, err := object.NewStatus(account_id, content)
	if err != nil {
		return nil, err
	}

	err = s.unitOfWork.Do(ctx, func(tx *sqlx.Tx) error {
		err = s.statusRepo.Create(ctx, tx, st)
		if err != nil {
			return err
		}
		return nil
	})

	return &CreateStatusDTO{
		Status: st,
	}, nil
}

func (s *status) FindByID(ctx context.Context, id string) (*GetStatusDTO, error) {
	st, err := s.statusRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &GetStatusDTO{
		Status: st,
	}, nil
}

func (s *status) FindPublicTimeline(ctx context.Context, limit int) ([]*object.Status, error) {
	st, err := s.statusRepo.FindPublicTimeline(ctx, limit)
	if err != nil {
		return nil, err
	}

	return st, nil
}
