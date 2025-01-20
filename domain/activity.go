package domain

type Activity struct {
	ActivityID        string  `db:"activity_id"`
	ActivityType      int     `db:"activity_type"`
	DurationInMinutes int     `db:"duration_in_minutes"`
	DoneAt            string  `db:"done_at"`
	CaloriesBurned    float64 `db:"calories_burned"`
	CreatedAt         string  `db:"created_at"`
	UpdatedAt         string  `db:"updated_at"`
}
