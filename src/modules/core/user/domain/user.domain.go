package userDomain

import "github.com/quinntas/go-gin-template/src/internal/ddd"

type User struct {
	ddd.BaseDomain
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Role     string `json:"role" db:"role"`
}

func NewUser(baseDomain ddd.BaseDomain, email string, password string, role string) User {
	return User{baseDomain, email, password, role}
}
