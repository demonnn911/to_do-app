package todo

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
