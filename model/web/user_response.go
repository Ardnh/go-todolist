package web

type UserResponseByUsername struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
	Password  string
}

type UserResponse struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
}

type ResponseWithToken struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
	Token     string
}

type UserResponseWithToken struct {
	Code   int
	Status string
	Data   ResponseWithToken
}
