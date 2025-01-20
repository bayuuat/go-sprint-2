package domain

type ActivityTypes struct {
	Id                int     `db:"id"`
	Name              string  `db:"name"`
	CaloriesPerMinute float64 `db:"calories_per_minute"`
}
