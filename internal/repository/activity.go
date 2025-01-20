package repository

import (
	"context"
	"database/sql"
	"time"
	"github.com/bayuuat/go-sprint-2/domain"
	"github.com/bayuuat/go-sprint-2/dto"
	"github.com/doug-martin/goqu/v9"
)

type ActivityRepository interface {
	Save(ctx context.Context, activity *domain.Activity) (*domain.Activity, error)
	Update(ctx context.Context, userId string, activity goqu.Record) error
	FindAllWithFilter(ctx context.Context, filter *dto.ActivityFilter) ([]domain.Activity, error)
	FindById(ctx context.Context, userId, id string) (domain.Activity, error)
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

func (d activityRepository) Update(ctx context.Context, userId string, activity goqu.Record) error {
	executor := d.db.Update("activities").Where(goqu.C("activity_id").Eq(userId)).Set(activity).Executor()
	_, err := executor.ExecContext(ctx)

	return err
}

func (d activityRepository) FindById(ctx context.Context, userId, id string) (activity domain.Activity, err error) {
	dataset := d.db.From("activities").Where(goqu.Ex{
		"activity_id": id,
	})
	_, err = dataset.ScanStructContext(ctx, &activity)
	return
}

func (r *activityRepository) HasEmployees(ctx context.Context, activityId string) (bool, error) {
	return false, nil
}

func (d activityRepository) Delete(ctx context.Context, user_id string, id string) error {
	return nil
}

func (d activityRepository) FindAllWithFilter(ctx context.Context, filter *dto.ActivityFilter) ([]domain.Activity, error) {
	query := d.db.From("activities")

	if filter.Limit > 0 {
		query = query.Limit(uint(filter.Limit))
	}
	if filter.Offset > 0 {
		query = query.Offset(uint(filter.Offset))
	}

	if filter.ActivityId != "" {
		query = query.Where(goqu.Ex{"activityid": filter.ActivityId})
	}

	if filter.ActivityType != "" {
		query = query.Where(goqu.Ex{"activitytype": filter.ActivityType})
	}

	if filter.DoneAtFrom != "" {
		if doneAtFrom, err := time.Parse(time.RFC3339, filter.DoneAtFrom); err == nil {
			query = query.Where(goqu.C("doneat").Gte(doneAtFrom))
		}
	}

	if filter.DoneAtTo != "" {
		if doneAtTo, err := time.Parse(time.RFC3339, filter.DoneAtTo); err == nil {
			query = query.Where(goqu.C("doneat").Lte(doneAtTo))
		}
	}

	if filter.CaloriesBurnedMin > 0 {
		query = query.Where(goqu.C("caloriesburned").Gte(filter.CaloriesBurnedMin))
	}

	if filter.CaloriesBurnedMax > 0 {
		query = query.Where(goqu.C("caloriesburned").Lte(filter.CaloriesBurnedMax))
	}

	query = query.Limit(uint(filter.Limit)).Offset(uint(filter.Offset))

	var activities []domain.Activity
	err := query.ScanStructsContext(ctx, &activities)
	return activities, err
}
