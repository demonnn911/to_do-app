package todo

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
	Email    string `json:"email" binding:"required"`
}
