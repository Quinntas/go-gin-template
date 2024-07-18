package main

import (
	"github.com/quinntas/go-gin-template/src/internal/api/cache/redis"
	"github.com/quinntas/go-gin-template/src/internal/api/database/mysql"
	"github.com/quinntas/go-gin-template/src/internal/api/web"
	"github.com/quinntas/go-gin-template/src/internal/common/envLoader"
	"github.com/quinntas/go-gin-template/src/modules/shared"
)

func main() {
	env, err := envLoader.NewVariables()
	if err != nil {
		panic(err)
	}

	rwDB, err := mysql.NewMysql(env.DatabaseURL)
	if err != nil {
		panic(err)
	}

	cache, err := redis.NewRedis(env.RedisURL)
	if err != nil {
		panic(err)
	}

	app := web.NewApp(
		env,
		rwDB,
		rwDB,
		cache,
	)

	shared.V1Router(app)

	err = app.Run()
	if err != nil {
		panic(err)
	}
}
