package main

import (
	"github.com/gin-gonic/gin"
	"github.com/quinntas/go-gin-template/src/internal/api/web"
)

func main() {
	r := gin.Default()

	web.Get(r, "/", func(c *gin.Context) error {
		web.JsonResponse(c, 200, gin.H{
			"message": "Hello, World!"})
		return nil
	})

	err := r.Run()

	if err != nil {
		return
	}
}
