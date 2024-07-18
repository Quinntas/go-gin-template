package web

import "github.com/gin-gonic/gin"

type Controller func(app *App, c *gin.Context) error
