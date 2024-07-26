package web

import "github.com/gin-gonic/gin"

func Route(routerGroup *gin.RouterGroup, method string, path string, controller Controller) {
	routerGroup.Handle(method, path, func(c *gin.Context) {
		handler(c, controller)
	})
}

func Get(routerGroup *gin.RouterGroup, path string, controller Controller) {
	Route(routerGroup, "GET", path, controller)
}

func Post(routerGroup *gin.RouterGroup, path string, controller Controller) {
	Route(routerGroup, "POST", path, controller)
}
