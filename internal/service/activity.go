package service

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/bayuuat/go-sprint-2/domain"
	"github.com/bayuuat/go-sprint-2/dto"
	"github.com/bayuuat/go-sprint-2/internal/config"
	"github.com/bayuuat/go-sprint-2/internal/repository"
	"github.com/google/uuid"
)

type activityService struct {
	cnf                     *config.Config
	activityRepository      repository.ActivityRepository
	activityTypesRepository repository.ActivityTypesRepository
}

type ActivityService interface {
	GetActivitysWithFilter(ctx context.Context, filter dto.ActivityFilter) ([]dto.ActivityData, int, error)
	CreateActivity(ctx context.Context, req dto.ActivityReq) (dto.ActivityData, int, error)
	PatchActivity(ctx context.Context, req dto.UpdateActivityReq, id, userId string) (dto.ActivityData, int, error)
	DeleteActivity(ctx context.Context, user_id string, id string) (dto.ActivityData, int, error)
}

func NewActivity(cnf *config.Config,
	activityRepository repository.ActivityRepository,
	activityTypesRepository repository.ActivityTypesRepository) ActivityService {
	return &activityService{
		cnf:                     cnf,
		activityRepository:      activityRepository,
		activityTypesRepository: activityTypesRepository,
	}
}

func (ds activityService) GetActivitysWithFilter(ctx context.Context, filter dto.ActivityFilter) ([]dto.ActivityData, int, error) {
	return []dto.ActivityData{}, http.StatusOK, nil
}

func (ds *activityService) CreateActivity(ctx context.Context, req dto.ActivityReq) (dto.ActivityData, int, error) {
	activityTypeId, err := strconv.Atoi(req.ActivityType)
	if err != nil {
		return dto.ActivityData{}, http.StatusBadRequest, err
	}

	activityType, err := ds.activityTypesRepository.FindById(ctx, activityTypeId)
	if err != nil || activityType.Id == 0 {
		return dto.ActivityData{}, http.StatusBadRequest, err
	}

	doneAt, err := time.Parse(time.RFC3339, req.DoneAt)
	if err != nil {
		return dto.ActivityData{}, http.StatusBadRequest, err
	}

	var createdAt, updatedAt sql.NullTime
	createdAt.Time = time.Now()
	updatedAt.Time = time.Now()
	createdAt.Valid = true
	updatedAt.Valid = true

	activity := domain.Activity{
		ActivityId:        uuid.New().String(),
		ActivityType:      activityTypeId,
		DoneAt:            doneAt,
		DurationInMinutes: req.DurationInMinutes,
		CreatedAt:         createdAt,
		UpdatedAt:         updatedAt,
	}

	success, err := ds.activityRepository.Save(ctx, &activity)
	if err != nil || success == nil {
		return dto.ActivityData{}, http.StatusInternalServerError, err
	}

	return dto.ActivityData{
		ActivityId:        activity.ActivityId,
		ActivityType:      strconv.Itoa(activity.ActivityType),
		DoneAt:            activity.DoneAt.Format(time.RFC3339),
		DurationInMinutes: activity.DurationInMinutes,
		CaloriesBurned:    activity.CaloriesBurned,
		CreatedAt:         activity.CreatedAt.Time.Format(time.RFC3339),
		UpdatedAt:         activity.UpdatedAt.Time.Format(time.RFC3339),
	}, http.StatusCreated, nil
}

func (ds activityService) PatchActivity(ctx context.Context, req dto.UpdateActivityReq, id, userId string) (dto.ActivityData, int, error) {
	return dto.ActivityData{}, http.StatusOK, nil
}

func (ds activityService) DeleteActivity(ctx context.Context, user_id string, id string) (dto.ActivityData, int, error) {
	return dto.ActivityData{}, http.StatusOK, nil
}
