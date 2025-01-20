package service

import (
	"context"
	"encoding/json"
	"github.com/bayuuat/go-sprint-2/domain"
	"github.com/bayuuat/go-sprint-2/dto"
	"github.com/bayuuat/go-sprint-2/internal/config"
	"github.com/bayuuat/go-sprint-2/internal/repository"
	"log/slog"
	"net/http"
)

type activityService struct {
	cnf                *config.Config
	activityRepository repository.ActivityRepository
}

type ActivityService interface {
	GetActivitiesWithFilter(ctx context.Context, filter dto.ActivityFilter) ([]dto.ActivityData, int, error)
	CreateActivity(ctx context.Context, req dto.ActivityReq, userId string) (dto.ActivityData, int, error)
	PatchActivity(ctx context.Context, req dto.UpdateActivityReq, userId, id string) (dto.ActivityData, int, error)
	DeleteActivity(ctx context.Context, id string) (dto.ActivityData, int, error)
}

func NewActivity(cnf *config.Config,
	activityRepository repository.ActivityRepository) ActivityService {
	return &activityService{
		cnf:                cnf,
		activityRepository: activityRepository,
	}
}

func (ds activityService) GetActivitiesWithFilter(ctx context.Context, filter dto.ActivityFilter) ([]dto.ActivityData, int, error) {
	return []dto.ActivityData{}, http.StatusOK, nil
}

func (ds activityService) CreateActivity(ctx context.Context, req dto.ActivityReq, UserId string) (dto.ActivityData, int, error) {
	return dto.ActivityData{}, http.StatusOK, nil
}

func (ds activityService) PatchActivity(ctx context.Context, req dto.UpdateActivityReq, userId, id string) (dto.ActivityData, int, error) {
	activity, err := ds.activityRepository.FindById(ctx, userId, id)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.ActivityData{}, http.StatusInternalServerError, err
	}

	if activity.ActivityID == "" {
		return dto.ActivityData{}, http.StatusNotFound, domain.ErrActivityNotFound
	}

	convertToActivity := map[string]int{
		"walking":    1,
		"yoga":       2,
		"stretching": 3,
		"cycling":    4,
		"swimming":   5,
		"dancing":    6,
		"hiking":     7,
		"running":    8,
		"hiit":       9,
		"jumprope":   10,
	}

	convertToActivityWord := [...]string{
		"",
		"walking",
		"yoga",
		"stretching",
		"cycling",
		"swimming",
		"dancing",
		"hiking",
		"running",
		"hiit",
		"jumprope",
	}

	reqDb := &dto.UpdateActivityDbReq{
		ActivityType:      convertToActivity[req.ActivityType],
		DoneAt:            req.DoneAt,
		DurationInMinutes: req.DurationInMinutes,
	}

	var activityMap map[string]interface{}

	jsonString, err := json.Marshal(&reqDb)
	if err != nil {
		return dto.ActivityData{}, http.StatusInternalServerError, err
	}

	err = json.Unmarshal(jsonString, &activityMap)
	if err != nil {
		return dto.ActivityData{}, http.StatusInternalServerError, err
	}

	if len(activityMap) == 0 {
		return dto.ActivityData{
			ActivityType:      convertToActivityWord[activity.ActivityType],
			DoneAt:            activity.DoneAt,
			DurationInMinutes: activity.DurationInMinutes,
			CaloriesBurned:    activity.CaloriesBurned,
			CreatedAt:         activity.CreatedAt,
			UpdatedAt:         activity.UpdatedAt,
		}, 200, nil
	}

	err = ds.activityRepository.Update(ctx, id, activityMap)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.ActivityData{}, http.StatusInternalServerError, err
	}

	// After update
	activity, err = ds.activityRepository.FindById(ctx, userId, id)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.ActivityData{}, http.StatusInternalServerError, err
	}

	return dto.ActivityData{
		ActivityType:      convertToActivityWord[activity.ActivityType],
		DoneAt:            activity.DoneAt,
		DurationInMinutes: activity.DurationInMinutes,
		CaloriesBurned:    activity.CaloriesBurned,
		CreatedAt:         activity.CreatedAt,
		UpdatedAt:         activity.UpdatedAt,
	}, 200, nil
}

func (ds activityService) DeleteActivity(ctx context.Context, id string) (dto.ActivityData, int, error) {
	return dto.ActivityData{}, http.StatusOK, nil
}
