package models

import (
	// "strconv"
	// "time"
	// "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


// DROP TABLE IF EXISTS  `auth_login`;
// CREATE TABLE `auth_login` (
//   `_id` char(32) NOT NULL COMMENT 'username,phone,email,wx_openid,...',
//   `salt` varchar(45) DEFAULT NULL,
//   `hash_pwd` varchar(45) DEFAULT NULL,
//   `account_id` char(32) DEFAULT NULL,
//   `ctime` bigint(19) NOT NULL DEFAULT '0',
//   PRIMARY KEY (`_id`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

func FindAuthLogin(uid string) *AuthLogin {
	//查询数据
	var salt string
	var hash_pwd string
  var account_id string
	err := GlobalMysqlConnPool.QueryRow("SELECT salt,hash_pwd,account_id FROM auth_login WHERE _id=?", uid).Scan(&salt,&hash_pwd,&account_id)
	if (err != nil) {
		return nil
	}

	var login = &AuthLogin{}
	login.Id = uid
	login.Salt = salt
	login.HashPwd = hash_pwd
	login.AccountId = account_id
	fmt.Println(login)

	return login
}


func AddAuthLogin(login AuthLogin) {
	//插入数据
	stmt, err := GlobalMysqlConnPool.Prepare("INSERT auth_login SET _id=?,salt=?,hash_pwd=?,account_id=?,ctime=?")
	checkErr(err)

	res, err := stmt.Exec(login.Id, login.Salt, login.HashPwd, login.AccountId, login.Ctime)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
}


func UpdateAuthLogin(uid string, salt string, hash_pwd string) {
	//更新数据
	stmt, err := GlobalMysqlConnPool.Prepare("UPDATE auth_login set salt=?,hash_pwd=? WHERE _id=?")
	checkErr(err)

	res, err := stmt.Exec(salt, hash_pwd, uid)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}
