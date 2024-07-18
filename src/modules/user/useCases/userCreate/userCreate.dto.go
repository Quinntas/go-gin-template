package userCreate

import "github.com/quinntas/go-gin-template/src/internal/api/web"

type DTO struct {
	App      *web.App
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseDTO struct {
	Message string `json:"message"`
}
