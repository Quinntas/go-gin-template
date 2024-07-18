package userDomain

import (
	"github.com/quinntas/go-gin-template/src/internal/ddd"
	"time"
)

type User struct {
	ddd.BaseDomain
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
}

func NewUser(id int, pid string, updatedAt time.Time, createdAt time.Time, email, password string, role string) User {
	return User{
		ddd.NewBaseDomain(id, pid, updatedAt, createdAt),
		email,
		password,
		role,
	}
}
