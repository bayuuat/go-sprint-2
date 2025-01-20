package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/bayuuat/go-sprint-2/domain"
	"github.com/doug-martin/goqu/v9"
)

type ActivityTypesRepository interface {
	Save(ctx context.Context, activityType *domain.ActivityTypes) (*domain.ActivityTypes, error)
	Update(ctx context.Context, activityType *domain.ActivityTypes) error
	FindAll(ctx context.Context) ([]domain.ActivityTypes, error)
	FindById(ctx context.Context, id int) (activityType domain.ActivityTypes, err error)
	Delete(ctx context.Context, id string) error
}

type activityTypesRepository struct {
	db *goqu.Database
}

func NewActivityType(db *sql.DB) ActivityTypesRepository {
	return &activityTypesRepository{
		db: goqu.New("default", db),
	}
}

func (a activityTypesRepository) Save(ctx context.Context, activityType *domain.ActivityTypes) (*domain.ActivityTypes, error) {
	return nil, nil
}

func (a activityTypesRepository) Update(ctx context.Context, activityType *domain.ActivityTypes) error {
	return nil
}

func (a activityTypesRepository) FindAll(ctx context.Context) ([]domain.ActivityTypes, error) {
	return nil, nil
}

func (a activityTypesRepository) FindById(ctx context.Context, id int) (activityType domain.ActivityTypes, err error) {
	dataset := a.db.From("activity_types").Where(goqu.Ex{
		"id": id,
	})

	_, err = dataset.ScanStructContext(ctx, &activityType)
	if err != nil {
		// Log the error
		fmt.Printf("Error finding activity type by id %d: %v\n", id, err)
	} else {
		// Log the successful retrieval
		fmt.Printf("Successfully found activity type by id %d: %+v\n", id, activityType)
	}

	return
}

func (a activityTypesRepository) Delete(ctx context.Context, id string) error {
	return nil
}
