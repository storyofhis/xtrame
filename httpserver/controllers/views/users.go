package views

import (
	"time"

	"github.com/storyofhis/xtrame/httpserver/repository/models"
)

type Register struct {
	Id        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	NickName  string    `json:"nick_name"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       uint      `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	Role      models.Role
}

type Login struct {
	Token string `json:"token"`
}
