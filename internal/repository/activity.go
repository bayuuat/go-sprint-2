package repository

import (
	"context"
	"database/sql"

	"github.com/bayuuat/go-sprint-2/domain"
	"github.com/bayuuat/go-sprint-2/dto"
	"github.com/doug-martin/goqu/v9"
)

type ActivityRepository interface {
	Save(ctx context.Context, activity *domain.Activity) (*domain.Activity, error)
	Update(ctx context.Context, activity *domain.Activity) error
	FindAllWithFilter(ctx context.Context, filter *dto.ActivityFilter) ([]domain.Activity, error)
	FindById(ctx context.Context, id string, userId string) (domain.Activity, error)
	HasEmployees(ctx context.Context, activityId string) (bool, error)
	Delete(ctx context.Context, user_id string, id string) error
}

type activityRepository struct {
	db *goqu.Database
}

func NewActivity(db *sql.DB) ActivityRepository {
	return &activityRepository{
		db: goqu.New("default", db),
	}
}

func (d activityRepository) Save(ctx context.Context, activity *domain.Activity) (*domain.Activity, error) {
	executor := d.db.Insert("activities").Rows(activity).Executor()
	_, err := executor.ExecContext(ctx)
	return activity, err
}

func (d activityRepository) Update(ctx context.Context, activity *domain.Activity) error {
	return nil
}

func (d activityRepository) FindById(ctx context.Context, id, userId string) (activity domain.Activity, err error) {
	return domain.Activity{}, nil
}

func (r *activityRepository) HasEmployees(ctx context.Context, activityId string) (bool, error) {
	return false, nil
}

func (d activityRepository) Delete(ctx context.Context, user_id string, id string) error {
	return nil
}

func (d activityRepository) FindAllWithFilter(ctx context.Context, filter *dto.ActivityFilter) ([]domain.Activity, error) {
	return nil, nil
}
