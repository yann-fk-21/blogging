package types

type UserStore interface {
	CreateUser(u UserRegister) error
	//GetUserByID(id int) (*User, error)
	//GetUsers()([]User, error)
	//UpdateUser(id int, u User) error
	//DeleteUser(id int) error
}

type UserRegister struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}
