package web

import (
	"github.com/gin-gonic/gin"
	"github.com/quinntas/go-gin-template/src/internal/api/cache"
	"github.com/quinntas/go-gin-template/src/internal/api/database"
	envLoader "github.com/quinntas/go-gin-template/src/internal/common/envLoader"
)

type App struct {
	Engine *gin.Engine
	Env    *envLoader.Variables
	RDB    database.Database
	WDB    database.Database
	Cache  cache.Cache
}

func NewApp(env *envLoader.Variables, rDB database.Database, wDB database.Database, cache cache.Cache) *App {
	engine := gin.Default()

	return &App{
		engine,
		env,
		rDB,
		wDB,
		cache,
	}
}

func (app *App) Run() error {
	return app.Engine.Run()
}
