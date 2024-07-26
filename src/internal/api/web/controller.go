package web

import "github.com/gin-gonic/gin"

type Controller func(c *gin.Context) error
