package createUser

import (
	"github.com/quinntas/go-gin-template/src/internal/api/web"
	"github.com/quinntas/go-gin-template/src/internal/utils"
)

func EmailAlreadyRegistered() *web.HttpError {
	return web.NewHttpError(409, "Email already registered", utils.H{})
}
