package domain

import (
	"database/sql"
	"time"
)

type Activity struct {
	ActivityId        string       `db:"activityid"`
	ActivityType      int          `db:"activitytype"`
	DoneAt            time.Time    `db:"doneat"`
	DurationInMinutes int          `db:"durationinminutes"`
	CaloriesBurned    float64      `db:"caloriesburned"`
	CreatedAt         sql.NullTime `db:"createdat"`
	UpdatedAt         sql.NullTime `db:"updatedat"`
}
