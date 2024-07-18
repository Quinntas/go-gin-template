package web

import "github.com/gin-gonic/gin"

func Route(routerGroup *gin.RouterGroup, app *App, method string, path string, controller Controller) {
	routerGroup.Handle(method, path, func(c *gin.Context) {
		handler(app, c, controller)
	})
}

func Get(routerGroup *gin.RouterGroup, app *App, path string, controller Controller) {
	Route(routerGroup, app, "GET", path, controller)
}

func Post(routerGroup *gin.RouterGroup, app *App, path string, controller Controller) {
	Route(routerGroup, app, "POST", path, controller)
}
