package dto

type ActivityReq struct {
	ActivityType      string `json:"activityType" validate:"required,oneof=1 2 3 4 5 6 7 8 9 10"`
	DoneAt            string `json:"doneAt" validate:"required,rfc3339"`
	DurationInMinutes int    `json:"durationInMinutes" validate:"required,min=1"`
}

type ActivityFilter struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Name   string `json:"name"`
	UserId string `json:"user_id"`
}

type ActivityData struct {
	ActivityId        string  `json:"activityId"`
	ActivityType      string  `json:"activityType"`
	DoneAt            string  `json:"doneAt"`
	DurationInMinutes int     `json:"durationInMinutes"`
	CaloriesBurned    float64 `json:"caloriesBurned"`
	CreatedAt         string  `json:"createdAt"`
	UpdatedAt         string  `json:"updatedAt"`
}

type UpdateActivityReq struct {
	Name string `json:"name" validate:"required,min=4,max=33"`
}
