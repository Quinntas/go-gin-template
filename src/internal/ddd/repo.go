package ddd

import (
	"github.com/quinntas/go-gin-template/src/internal/api/cache"
	"github.com/quinntas/go-gin-template/src/internal/api/database"
)

type Repo struct {
	readConn  *database.Database
	writeConn *database.Database
	cache     *cache.Cache
}
