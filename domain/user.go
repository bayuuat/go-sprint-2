package domain

import (
	"time"
)

type User struct {
	Id         string    `db:"id"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	Preference string    `db:"preference"`
	WeightUnit string    `db:"weight_unit"`
	HeightUnit string    `db:"height_unit"`
	Weight     float64   `db:"weight"`
	Height     float64   `db:"height"`
	Name       *string   `db:"name"`
	ImageURI   *string   `db:"image_uri"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
