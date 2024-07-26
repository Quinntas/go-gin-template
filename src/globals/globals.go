package globals

import (
	"github.com/quinntas/go-gin-template/src/internal/api/cache/redis"
	"github.com/quinntas/go-gin-template/src/internal/api/database/mysql"
	"github.com/quinntas/go-gin-template/src/internal/common/envLoader"
)

var Env *envLoader.Variables
var RWDatabase *mysql.Mysql
var Cache *redis.Redis
