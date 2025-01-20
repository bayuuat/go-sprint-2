package dto

type ActivityReq struct {
	ActivityType      string `json:"activityType" validate:"required,oneof=Walking Yoga Stretching Cycling Swimming Dancing Hiking Running HIIT JumpRope"`
	DoneAt            string `json:"doneAt" validate:"required,rfc3339"`
	DurationInMinutes int    `json:"durationInMinutes" validate:"required,min=1"`
}

type ActivityFilter struct {
	Limit             int    `json:"limit"`
	Offset            int    `json:"offset"`
	ActivityId        string `json:"activityId"`
	ActivityType      string `json:"activityType"`
	DoneAtFrom        string `json:"doneAtFrom"`
	DoneAtTo          string `json:"doneAtTo"`
	CaloriesBurnedMin int    `json:"caloriesBurnedMin"`
	CaloriesBurnedMax int    `json:"caloriesBurnedMax"`
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
