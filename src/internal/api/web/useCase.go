package web

import "github.com/gin-gonic/gin"

type UseCase func(c *gin.Context) error
