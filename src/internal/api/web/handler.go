package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, err error) {
	switch t := err.(type) {
	case *HttpError:
		t.serialize(c)
	default:
		InternalServerError().serialize(c)
	}
}

func handler(c *gin.Context, controller Controller) {
	err := controller(c)
	if err != nil {
		handleError(c, err)
	}
	fmt.Println(err.Error())
}
