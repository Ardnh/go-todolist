package web

type UserResponse struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
}

type UserResponseWithToken struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
	Token     string
}
