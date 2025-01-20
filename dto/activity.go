package dto

type ActivityReq struct {
	ActivityType      string `json:"activityType" validate:"required,oneof=walking yoga stretching cycling swimming dancing hiking running hiit jumprope"`
	DoneAt            string `json:"doneAt" validate:"required,isodate"`
	DurationInMinutes int    `json:"durationInMinutes" validate:"required,min=1"`
}

type ActivityFilter struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Name   string `json:"name"`
	UserId string `json:"user_id"`
}

type ActivityData struct {
	ActivityType      string  `json:"activityType"`
	DoneAt            string  `json:"doneAt"`
	DurationInMinutes int     `json:"durationInMinutes"`
	CaloriesBurned    float64 `json:"caloriesBurned"`
	CreatedAt         string  `json:"createdAt"`
	UpdatedAt         string  `json:"updatedAt"`
}

type UpdateActivityReq struct {
	ActivityType      string `json:"activityType,omitempty" validate:"omitempty,oneof=walking yoga stretching cycling swimming dancing hiking running hiit jumprope"`
	DoneAt            string `json:"doneAt,omitempty" validate:"omitempty,isodate"`
	DurationInMinutes int    `json:"durationInMinutes,omitempty" validate:"omitempty,min=0"`
}

type UpdateActivityDbReq struct {
	ActivityType      int    `json:"activity_type,omitempty" validate:"omitempty,oneof=walking yoga stretching cycling swimming dancing hiking running hiit jumprope"`
	DoneAt            string `json:"done_at,omitempty" validate:"omitempty,isodate"`
	DurationInMinutes int    `json:"duration_in_minutes,omitempty" validate:"omitempty,min=0"`
}
