package model

import "github.com/Zhoangp/Api_Gateway-Courses/pkg/common"

type ChangePassRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}
type NewInstrcutorRequest struct {
	Website  string `json:"website"`
	LinkedIn string `json:"linkedin"`
	Youtube  string `json:"youtube"`
	Bio      string `json:"bio"`
}
type GetTokenRequest struct {
	Email string `json:"email"`
}
type UpdateUserRequest struct {
	common.SQLModel
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}
