package service

import (
	"context"
	"net/http"
	"strconv"
	"time"

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
	activities, err := ds.activityRepository.FindAllWithFilter(ctx, &filter)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	if len(activities) == 0 {
		return []dto.ActivityData{}, http.StatusOK, nil
	}

	var activityData []dto.ActivityData
	for _, v := range activities {
		activityData = append(activityData, dto.ActivityData{
			ActivityId:        v.ActivityId,
			ActivityType:      strconv.Itoa(v.ActivityType),
			DoneAt:            v.DoneAt.Format(time.RFC3339),
			DurationInMinutes: v.DurationInMinutes,
			CaloriesBurned:    v.CaloriesBurned,
			CreatedAt:         v.CreatedAt.Time.Format(time.RFC3339),
		})
	}
	return activityData, http.StatusOK, nil
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
