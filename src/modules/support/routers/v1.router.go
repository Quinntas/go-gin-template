package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/quinntas/go-gin-template/src/modules/core/user/infra/http"
)

func V1(engine *gin.Engine) {
	versionRouter := engine.Group("/api/v1")

	userRouter.InitRouter(versionRouter)
}
