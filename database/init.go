package database

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
)

var dbConnect *pg.DB

func init() {
	fmt.Println("connect db.....")

	dbConnect = pg.Connect(&pg.Options{
		Addr:        "0.0.0.0:5432",
		User:        "postgres",
		Password:    "123123",
		Database:    "products",
		MaxRetries:  3,
		DialTimeout: 3 * time.Second, //	mặc định là 5
		ReadTimeout: 1 * time.Second,
		// Timeout for socket writes. If reached, commands will fail
		// with a timeout instead of blocking.
		WriteTimeout: 1 * time.Second,

		// Default is 10 connections per every CPU as reported by runtime.NumCPU.
		PoolSize: 5,
	})
}

// GetConn return the db connection
func GetConn() (db *pg.DB) {

	return dbConnect
}
