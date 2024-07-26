package mysql

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/quinntas/go-gin-template/src/internal/api/database"
	"github.com/quinntas/go-gin-template/src/internal/utils"
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

func (db *Mysql) Insert(query string, args ...interface{}) (database.InsertResult, error) {
	res, err := db.client.Exec(query, args...)
	if err != nil {
		return database.InsertResult{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return database.InsertResult{}, err
	}
	return database.NewInsertResult(id), nil
}

func (db *Mysql) Update(query string, args ...interface{}) (database.UpdateResult, error) {
	res, err := db.client.Exec(query, args...)
	if err != nil {
		return database.UpdateResult{}, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return database.UpdateResult{}, err
	}
	return database.NewUpdateResult(rows), nil
}

func (db *Mysql) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.client.Exec(query, args...)
}

func (db *Mysql) Query(query string, args ...interface{}) ([]utils.H, error) {
	rows, err := db.client.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]utils.H, 0)
	for rows.Next() {
		var row utils.H
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
