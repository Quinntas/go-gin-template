package web

import "github.com/gin-gonic/gin"

func JsonResponse[T gin.H](c *gin.Context, status int, body T) {
	c.JSON(status, body)
}
