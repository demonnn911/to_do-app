package todo

// while testing grpc authentication service, remove binding:required for username and email
type User struct {
	Id       int64  `json:"-" db:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type SignInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
