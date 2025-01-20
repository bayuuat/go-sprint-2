package domain

import "time"

type Activity struct {
	ActivityId        string    `db:"activity_id"`
	ActivityType      int       `db:"activity_type"`
	DurationInMinutes int       `db:"duration_in_minutes"`
	DoneAt            time.Time `db:"done_at"`
	CaloriesBurned    float64   `db:"calories_burned"`
	CreatedAt         time.Time `db:"created_at"`
	UpdatedAt         time.Time `db:"updated_at"`
}
