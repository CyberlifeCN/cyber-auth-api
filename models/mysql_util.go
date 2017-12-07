package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//global
var GlobalMysqlConnPool *sql.DB


func init() {
	// init mysql connection pool
	GlobalMysqlConnPool, _ = sql.Open("mysql", "legend_dev:need4sPeed@tcp(rm-2zeyubz4yre340644o.mysql.rds.aliyuncs.com:3306)/legend_dev?charset=utf8mb4")
	GlobalMysqlConnPool.SetMaxOpenConns(20)
	GlobalMysqlConnPool.SetMaxIdleConns(10)
	GlobalMysqlConnPool.Ping()
	// defer GlobalMysqlConnPool.Close()
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
