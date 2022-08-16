/*
 * @Author: GG
 * @Date: 2022-08-16 11:55:08
 * @LastEditTime: 2022-08-16 11:59:35
 * @LastEditors: GG
 * @Description:mysql update
 * @FilePath: \golang-demo\database\mysql\curd\update.go
 *
 */
package curd

import (
	"database/sql"
	"fmt"
)

func Update(db *sql.DB) {
	sqlStr := "update user_tbl set username=?, password=? where id = ?"
	ret, err := db.Exec(sqlStr, "tom4", "169", 4)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	rows, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Println("rows:", rows)
}
