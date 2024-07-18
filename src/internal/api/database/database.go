package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var client *sqlx.DB

func Init(uri string) error {
	db, err := sqlx.Connect("mysql", uri)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	client = db
	return nil
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return client.Exec(query, args...)
}

func MustExec(query string, args ...interface{}) sql.Result {
	return client.MustExec(query, args...)
}

func Query[T interface{}](query string, args ...interface{}) ([]T, error) {
	rows, err := client.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	results := make([]T, 0)
	for rows.Next() {
		var row T
		err = rows.StructScan(&row)
		if err != nil {
			return nil, err
		}
		results = append(results, row)
	}
	return results, nil
}

func QueryRow[T interface{}](query string, args ...interface{}) (*T, error) {
	var result T
	err := client.QueryRowx(query, args...).StructScan(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
