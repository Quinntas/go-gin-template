package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/quinntas/go-gin-template/src/internal/utils"
)

type Database interface {
	Insert(query string, args ...interface{}) (InsertResult, error)
	Update(query string, args ...interface{}) (UpdateResult, error)
	Query(query string, args ...interface{}) ([]utils.H, error)
	Close() error
}
