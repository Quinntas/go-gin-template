package userRouter

import (
	"github.com/gin-gonic/gin"
	"github.com/quinntas/go-gin-template/src/internal/api/web"
	"github.com/quinntas/go-gin-template/src/modules/core/user/useCases/createUser"
)

func InitRouter(versionRouter *gin.RouterGroup) {
	userRouter := versionRouter.Group("/user")

	web.Post(userRouter, "/", createUser.Controller)
}
