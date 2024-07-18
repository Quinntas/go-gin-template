package database

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	MustExec(query string, args ...interface{}) sql.Result
	Query(query string, args ...interface{}) ([]gin.H, error)
}

type Client struct {
	client *sqlx.DB
}

func NewClient(uri string) (*Client, error) {
	conn, err := sqlx.Connect("mysql", uri)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return &Client{client: conn}, nil
}

func (db *Client) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.client.Exec(query, args...)
}

func (db *Client) MustExec(query string, args ...interface{}) sql.Result {
	return db.client.MustExec(query, args...)
}

func (db *Client) Query(query string, args ...interface{}) ([]gin.H, error) {
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
