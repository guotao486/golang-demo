/*
 * @Author: GG
 * @Date: 2022-08-16 09:59:50
 * @LastEditTime: 2022-08-16 11:15:47
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\database\mysql\select\select.go
 *
 */
package query

import (
	"database/sql"
	"fmt"
)

type User struct {
	id       int
	username string
	password string
}

func QueryRowDemo(DB *sql.DB) {
	sqlStr := "select id, username, password from user_tbl where id=?"
	var u User
	// 确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := DB.QueryRow(sqlStr, 1).Scan(&u.id, &u.username, &u.password)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%s\n", u.id, u.username, u.password)
}

func QueryMuiltDemo(DB *sql.DB) {
	sqlStr := "select id, username, password from user_tbl where id>?"
	// var u User
	rows, err := DB.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	defer rows.Close()

	// 依次处理
	for rows.Next() {
		var u User
		err := rows.Scan(&u.id, &u.username, &u.password)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		fmt.Printf("u: %v\n", u)
	}
}
