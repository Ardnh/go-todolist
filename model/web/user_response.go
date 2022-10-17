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
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	UserName  string `json:"username"`
	Token     string `json:"token"`
}

type UserResponseWithToken struct {
	Code   int               `json:"code"`
	Status string            `json:"status"`
	Data   ResponseWithToken `json:"data"`
}
