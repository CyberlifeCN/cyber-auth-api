package models

import (
	// "strconv"
	// "time"
	// "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


// DROP TABLE IF EXISTS  `auth_ticket`;
// CREATE TABLE `auth_ticket` (
//   `refresh_token` char(32) NOT NULL,
//   `access_token` char(32) DEFAULT NULL,
//   `expires_at` bigint(19) NOT NULL DEFAULT '0',
//   `account_id` char(32) DEFAULT NULL,
//   `token_type` varchar(45) DEFAULT 'Bearer',
//   `scope` varchar(45) DEFAULT 'ticket',
//   PRIMARY KEY (`refresh_token`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

func FindRefreshTicket(token string) *RefreshTicket {
	//查询数据
	var refresh_token string
	var access_token string
	var account_id string
	var expires_at int64
  var token_type string
  var scope string
	err := GlobalMysqlConnPool.QueryRow("SELECT refresh_token,access_token,account_id,expires_at,token_type,scope FROM auth_ticket WHERE refresh_token=?", token).Scan(&refresh_token,&access_token,&account_id,&expires_at,&token_type,&scope)
	if (err != nil) {
		return nil
	}

	var ticket = &RefreshTicket{}
	ticket.Id = refresh_token
	ticket.AccessToken = access_token
	ticket.AccountId = account_id
	ticket.ExpiresAt = expires_at
  ticket.TokenType = token_type
  ticket.Scope = scope
	fmt.Println(ticket)

	return ticket
}


func AddRefreshTicket(ticket RefreshTicket) {
	//插入数据
	stmt, err := GlobalMysqlConnPool.Prepare("INSERT auth_ticket SET refresh_token=?,access_token=?,account_id=?,expires_at=?")
	checkErr(err)

	res, err := stmt.Exec(ticket.Id, ticket.AccessToken, ticket.AccountId, ticket.ExpiresAt)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
}


func UpdateRefreshTicket(refresh_token string, access_token string) {
	//更新数据
	stmt, err := GlobalMysqlConnPool.Prepare("UPDATE auth_ticket set access_token=? WHERE refresh_token=?")
	checkErr(err)

	res, err := stmt.Exec(access_token, refresh_token)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}


func DeleteRefreshTicket(token string) {
	//删除数据
	stmt, err := GlobalMysqlConnPool.Prepare("DELETE FROM auth_ticket WHERE refresh_token=?")
	checkErr(err)

	res, err := stmt.Exec(token)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}
