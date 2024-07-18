package main

import (
	"github.com/gin-gonic/gin"
	"github.com/quinntas/go-gin-template/src/internal/api/web"
)

func main() {
	r := gin.Default()
	web.Get(r, "/ping", func(c *gin.Context) error {
		web.JsonResponse(c, 200, gin.H{
			"message": "pong"})
		return nil
	})
	err := r.Run()
	if err != nil {
		return
	}
}
