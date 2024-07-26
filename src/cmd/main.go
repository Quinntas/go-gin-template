package main

import (
	"github.com/gin-gonic/gin"
	"github.com/quinntas/go-gin-template/src/globals"
	"github.com/quinntas/go-gin-template/src/internal/api/cache/redis"
	"github.com/quinntas/go-gin-template/src/internal/api/database/mysql"
	"github.com/quinntas/go-gin-template/src/internal/common/envLoader"
	"github.com/quinntas/go-gin-template/src/modules/support/routers"
)

func main() {
	env, err := envLoader.NewVariables()
	if err != nil {
		panic(err)
	}

	globals.RWDatabase, err = mysql.NewMysql(env.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer globals.RWDatabase.Close()

	globals.Cache, err = redis.NewRedis(env.RedisURL)
	if err != nil {
		panic(err)
	}
	defer globals.Cache.Close()

	engine := gin.Default()

	routers.V1(engine)

	err = engine.Run()
	if err != nil {
		panic(err)
	}
}
