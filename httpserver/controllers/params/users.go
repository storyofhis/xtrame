package params

type Register struct {
	FullName string `json:"full_name" validate:"required"`
	NickName string `json:"nick_name" validate:"required"`
	UserName string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Age      uint   `json:"age" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
}

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
