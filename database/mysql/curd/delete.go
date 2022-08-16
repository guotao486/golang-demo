/*
 * @Author: GG
 * @Date: 2022-08-16 11:46:43
 * @LastEditTime: 2022-08-16 11:53:39
 * @LastEditors: GG
 * @Description:mysql delete
 * @FilePath: \golang-demo\database\mysql\curd\delete.go
 *
 */
package curd

import (
	"database/sql"
	"fmt"
)

func Delete(db *sql.DB) {
	sqlStr := "delete from user_tbl where id = ?"
	ret, err := db.Exec(sqlStr, 3)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	// 受到影响的行
	rows, err := ret.RowsAffected()
	fmt.Printf("rows: %v\n", rows)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("rows: %v\n", rows)

}
