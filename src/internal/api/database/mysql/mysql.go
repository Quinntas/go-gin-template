package mysql

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type Mysql struct {
	client *sqlx.DB
}

func NewMysql(uri string) (*Mysql, error) {
	conn, err := sqlx.Connect("mysql", uri)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return &Mysql{client: conn}, nil
}

func (db *Mysql) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.client.Exec(query, args...)
}

func (db *Mysql) Query(query string, args ...interface{}) ([]gin.H, error) {
	rows, err := db.client.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]gin.H, 0)
	for rows.Next() {
		var row gin.H
		err = rows.StructScan(&row)
		if err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	return results, nil
}

func (db *Mysql) Close() error {
	return db.client.Close()
}
