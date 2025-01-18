package dto

type AuthResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type AuthReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type UserPreferences struct {
	Preference *string `json:"preference"`
	WeightUnit *string `json:"weightUnit"`
	HeightUnit *string `json:"heightUnit"`
	Weight     *int    `json:"weight"`
	Height     *int    `json:"height"`
	Email      *string `json:"email"`
	Name       *string `json:"name"`
	ImageUri   *string `json:"imageUri"`
}

type UpdateUserPreferences struct {
	Preference *string  `json:"preference" validate:"required,oneof=CARDIO WEIGHT"`
	WeightUnit *string  `json:"weightUnit" validate:"required,oneof=KG LBS"`
	HeightUnit *string  `json:"heightUnit" validate:"required,oneof=CM INCH"`
	Weight     *float64 `json:"weight" validate:"required,min=10,max=1000"`
	Height     *float64 `json:"height" validate:"required,min=3,max=250"`
	Name       *string  `json:"name" validate:"required,min=2,max=60"`
	ImageUri   *string  `json:"imageUri" validate:"required,uri,accessibleuri"`
}
