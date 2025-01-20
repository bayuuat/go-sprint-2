package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/bayuuat/go-sprint-2/domain"
	"github.com/bayuuat/go-sprint-2/dto"
	"github.com/bayuuat/go-sprint-2/internal/config"
	"github.com/bayuuat/go-sprint-2/internal/repository"
	"github.com/google/uuid"
	"log/slog"
	"net/http"
	"time"
)

var activityTypeMap = map[string]int{
	"Walking":    1,
	"Yoga":       2,
	"Stretching": 3,
	"Cycling":    4,
	"Swimming":   5,
	"Dancing":    6,
	"Hiking":     7,
	"Running":    8,
	"HIIT":       9,
	"JumpRope":   10,
}

var convertToActivity = map[string]int{
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

var convertToActivityWord = [...]string{
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

type activityService struct {
	cnf                     *config.Config
	activityRepository      repository.ActivityRepository
	activityTypesRepository repository.ActivityTypesRepository
}

type ActivityService interface {
	GetActivitysWithFilter(ctx context.Context, filter dto.ActivityFilter) ([]dto.ActivityData, int, error)
	CreateActivity(ctx context.Context, req dto.ActivityReq, userId string) (dto.ActivityData, int, error)
	PatchActivity(ctx context.Context, req dto.UpdateActivityReq, userId, id string) (dto.ActivityData, int, error)
	DeleteActivity(ctx context.Context, user_id, id string) (dto.ActivityData, int, error)
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
			ActivityType:      convertToActivityWord[v.ActivityType],
			DoneAt:            v.DoneAt.Format(time.RFC3339),
			DurationInMinutes: v.DurationInMinutes,
			CaloriesBurned:    v.CaloriesBurned,
			CreatedAt:         v.CreatedAt.Format(time.RFC3339),
		})
	}
	return activityData, http.StatusOK, nil
}

func (ds *activityService) CreateActivity(ctx context.Context, req dto.ActivityReq, userId string) (dto.ActivityData, int, error) {
	activityTypeID, exists := activityTypeMap[req.ActivityType]
	if !exists {
		return dto.ActivityData{}, http.StatusBadRequest, errors.New("Not found")
	}

	activityType, err := ds.activityTypesRepository.FindById(ctx, activityTypeID)
	if err != nil || activityType.Id == 0 {
		return dto.ActivityData{}, http.StatusBadRequest, err
	}

	doneAt, err := time.Parse(time.RFC3339, req.DoneAt)
	if err != nil {
		return dto.ActivityData{}, http.StatusBadRequest, err
	}

	var createdAt, updatedAt time.Time
	createdAt = time.Now()
	updatedAt = time.Now()

	activity := domain.Activity{
		ActivityId:        uuid.New().String(),
		ActivityType:      activityTypeID,
		DoneAt:            doneAt,
		DurationInMinutes: req.DurationInMinutes,
		CreatedAt:         createdAt,
		UpdatedAt:         updatedAt,
	}

	success, err := ds.activityRepository.Save(ctx, &activity)
	if err != nil || success == nil {
		return dto.ActivityData{}, http.StatusInternalServerError, err
	}

	newActivity, err := ds.activityRepository.FindById(ctx, userId, activity.ActivityId)
	if err != nil && err != sql.ErrNoRows {
		return dto.ActivityData{}, http.StatusInternalServerError, err
	}

	return dto.ActivityData{
		ActivityId:        activity.ActivityId,
		ActivityType:      convertToActivityWord[activity.ActivityType],
		DoneAt:            activity.DoneAt.Format(time.RFC3339),
		DurationInMinutes: activity.DurationInMinutes,
		CaloriesBurned:    newActivity.CaloriesBurned,
		CreatedAt:         activity.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         activity.UpdatedAt.Format(time.RFC3339),
	}, http.StatusCreated, nil
}

func (ds activityService) PatchActivity(ctx context.Context, req dto.UpdateActivityReq, userId, id string) (dto.ActivityData, int, error) {
	activity, err := ds.activityRepository.FindById(ctx, userId, id)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.ActivityData{}, http.StatusInternalServerError, err
	}

	if activity.ActivityId == "" {
		return dto.ActivityData{}, http.StatusNotFound, domain.ErrActivityNotFound
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
			ActivityId:        activity.ActivityId,
			ActivityType:      convertToActivityWord[activity.ActivityType],
			DoneAt:            activity.DoneAt.Format(time.RFC3339),
			DurationInMinutes: activity.DurationInMinutes,
			CaloriesBurned:    activity.CaloriesBurned,
			CreatedAt:         activity.CreatedAt.Format(time.RFC3339),
			UpdatedAt:         activity.UpdatedAt.Format(time.RFC3339),
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
		ActivityId:        activity.ActivityId,
		ActivityType:      convertToActivityWord[activity.ActivityType],
		DoneAt:            activity.DoneAt.Format(time.RFC3339),
		DurationInMinutes: activity.DurationInMinutes,
		CaloriesBurned:    activity.CaloriesBurned,
		CreatedAt:         activity.CreatedAt.Format(time.RFC3339),
		UpdatedAt:         activity.UpdatedAt.Format(time.RFC3339),
	}, 200, nil
}

func (ds activityService) DeleteActivity(ctx context.Context, user_id, id string) (dto.ActivityData, int, error) {
	return dto.ActivityData{}, http.StatusOK, nil
}
