/*
 * @Author: GG
 * @Date: 2022-08-16 11:17:33
 * @LastEditTime: 2022-08-16 11:31:10
 * @LastEditors: GG
 * @Description: mysql insert
 * @FilePath: \golang-demo\database\mysql\curd\insert.go
 *
 */
package curd

import (
	"database/sql"
	"fmt"
)

// 插入 修改 删除 都使用exec
//func (db *DB) Exec(query string, args ...interface{}) (Result, error)

func InsertData(db *sql.DB) {
	sqlStr := "insert into user_tbl( username, password) value(?,?)"
	ret, err := db.Exec(sqlStr, "tom", "666")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	theId, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("theId: %v\n", theId)
}
