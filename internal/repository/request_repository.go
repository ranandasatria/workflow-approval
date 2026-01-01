package repository

import (
	"context"
	"workflow-approval/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RequestRepository interface {
	FindByIDForUpdate(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*model.Request, error)
	Update(ctx context.Context, tx *gorm.DB, req *model.Request) error
	Create(ctx context.Context, req *model.Request) error
    FindAll(ctx context.Context) ([]model.Request, error)
}

type requestRepo struct {
	db *gorm.DB
}

func NewRequestRepository(db *gorm.DB) RequestRepository {
	return &requestRepo{db}
}

func (r *requestRepo) FindByIDForUpdate(ctx context.Context, tx *gorm.DB, id uuid.UUID) (*model.Request, error) {
	var req model.Request
	err := tx.WithContext(ctx).Clauses(clause.Locking{Strength: "UPDATE"}).First(&req, "id = ?", id).Error
	return &req, err
}

func (r *requestRepo) Update(ctx context.Context, tx *gorm.DB, req *model.Request) error {
	return tx.WithContext(ctx).Save(req).Error
}

func (r *requestRepo) Create(ctx context.Context, req *model.Request) error {
    return r.db.WithContext(ctx).Create(req).Error
}

func (r *requestRepo) FindAll(ctx context.Context) ([]model.Request, error) {
    var requests []model.Request
    err := r.db.WithContext(ctx).Order("created_at desc").Find(&requests).Error
    return requests, err
}