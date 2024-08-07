package createUser

import (
	"github.com/gin-gonic/gin"
	"github.com/quinntas/go-gin-template/src/internal/api/web"
)

func Controller(c *gin.Context) error {
	var dto DTO

	if err := c.BindJSON(&dto); err != nil {
		return web.InvalidRequest()
	}

	res, err := UseCase(dto)

	if err != nil {
		return err
	}

	web.JsonResponse(c, 201, res)

	return nil
}
