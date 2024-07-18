package http

import (
	"github.com/gin-gonic/gin"
	"github.com/quinntas/go-gin-template/src/internal/api/web"
	"github.com/quinntas/go-gin-template/src/modules/user/useCases/userCreate"
)

func Router(router *gin.RouterGroup, app *web.App) {
	userRouter := router.Group("/users")

	web.Post(userRouter, app, "/", userCreate.Controller)
}
