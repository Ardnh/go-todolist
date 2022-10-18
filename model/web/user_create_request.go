package web

type CreateUserRequest struct {
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	UserName  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
