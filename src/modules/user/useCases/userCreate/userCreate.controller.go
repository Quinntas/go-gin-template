package userCreate

import (
	"github.com/gin-gonic/gin"
	"github.com/quinntas/go-gin-template/src/internal/api/web"
)

func Controller(app *web.App, c *gin.Context) error {
	dto := DTO{
		App: app,
	}

	err := c.BindJSON(&dto)
	if err != nil {
		return err
	}

	err = UseCase(dto)

	if err != nil {
		return err
	}

	web.JsonResponse[ResponseDTO](
		c,
		201,
		ResponseDTO{
			Message: "User created successfully",
		},
	)

	return nil
}
