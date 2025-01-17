package domain

import (
	"database/sql"
	"time"
)

type Activity struct {
	ActivityID        string       `db:"activity_id"`
	ActivityType      int          `db:"activity_type"`
	DoneAt            time.Time    `db:"done_at"`
	DurationInMinutes int          `db:"duration_in_minutes"`
	CaloriesBurned    float64      `db:"calories_burned"`
	CreatedAt         sql.NullTime `db:"created_at"`
	UpdatedAt         sql.NullTime `db:"updated_at"`
}
