package user

import (
	"github.com/quinntas/go-gin-template/src/internal/ddd"
	"time"
)

type User struct {
	ddd.BaseDomain
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

func NewUser(id int, pid string, updatedAt time.Time, createdAt time.Time, email, password string) User {
	return User{
		BaseDomain: ddd.NewBaseDomain(id, pid, updatedAt, createdAt),
		Email:      email,
		Password:   password,
	}
}
