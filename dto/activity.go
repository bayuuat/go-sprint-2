package dto

type ActivityReq struct {
	Name string `json:"name" validate:"required,min=4,max=33"`
}

type ActivityFilter struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Name   string `json:"name"`
	UserId string `json:"user_id"`
}

type ActivityData struct {
	DepartmentId string `json:"departmentId"`
	Name         string `json:"name"`
	UserId       string `json:"userId"`
}

type UpdateActivityReq struct {
	Name string `json:"name" validate:"required,min=4,max=33"`
}
