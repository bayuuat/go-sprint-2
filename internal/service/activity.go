package service

import (
	"context"
	"net/http"

	"github.com/bayuuat/go-sprint-2/dto"
	"github.com/bayuuat/go-sprint-2/internal/config"
	"github.com/bayuuat/go-sprint-2/internal/repository"
)

type activityService struct {
	cnf                *config.Config
	activityRepository repository.ActivityRepository
}

type ActivityService interface {
	GetActivitysWithFilter(ctx context.Context, filter dto.ActivityFilter) ([]dto.ActivityData, int, error)
	CreateActivity(ctx context.Context, req dto.ActivityReq, userId string) (dto.ActivityData, int, error)
	PatchActivity(ctx context.Context, req dto.UpdateActivityReq, id, userId string) (dto.ActivityData, int, error)
	DeleteActivity(ctx context.Context, user_id string, id string) (dto.ActivityData, int, error)
}

func NewActivity(cnf *config.Config,
	activityRepository repository.ActivityRepository) ActivityService {
	return &activityService{
		cnf:                cnf,
		activityRepository: activityRepository,
	}
}

func (ds activityService) GetActivitysWithFilter(ctx context.Context, filter dto.ActivityFilter) ([]dto.ActivityData, int, error) {
	return []dto.ActivityData{}, http.StatusOK, nil
}

func (ds activityService) CreateActivity(ctx context.Context, req dto.ActivityReq, UserId string) (dto.ActivityData, int, error) {
	return dto.ActivityData{}, http.StatusOK, nil
}

func (ds activityService) PatchActivity(ctx context.Context, req dto.UpdateActivityReq, id, userId string) (dto.ActivityData, int, error) {
	return dto.ActivityData{}, http.StatusOK, nil
}

func (ds activityService) DeleteActivity(ctx context.Context, user_id string, id string) (dto.ActivityData, int, error) {
	return dto.ActivityData{}, http.StatusOK, nil
}
