package web

import "github.com/gin-gonic/gin"

func Route(r *gin.Engine, method string, path string, controller Controller) {
	r.Handle(method, path, func(c *gin.Context) {
		handler(c, controller)
	})
}

func Get(r *gin.Engine, path string, controller Controller) {
	Route(r, "GET", path, controller)
}

func Post(r *gin.Engine, path string, controller Controller) {
	Route(r, "POST", path, controller)
}
