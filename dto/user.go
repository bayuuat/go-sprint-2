package dto

type AuthResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type AuthReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type UserData struct {
	Id              string  `json:"id"`
	Email           string  `json:"email"`
	Name            string  `json:"name"`
	UserImageUri    *string `json:"userImageUri"`
	CompanyName     *string `json:"companyName"`
	CompanyImageUri *string `json:"companyImageUri"`
}

type UpdateUserReq struct {
	Name            *string `json:"name" validate:"required,min=4,max=52"`
	Email           *string `json:"email" validate:"required,email"`
	UserImageUri    *string `json:"userImageUri" validate:"required,uri,accessibleuri"`
	CompanyName     *string `json:"companyName" validate:"required,min=4,max=52"`
	CompanyImageUri *string `json:"companyImageUri" validate:"required,uri,accessibleuri"`
}
