/*
 * @Author: GG
 * @Date: 2022-08-18 15:05:31
 * @LastEditTime: 2022-08-18 18:11:43
 * @LastEditors: GG
 * @Description:user dao
 * @FilePath: \golang-demo\blog\golang\dao\user.go
 *
 */
package dao

import (
	"golang-demo/blog/golang/models"
	"log"
)

func GetUserNameById(id int) (name string) {
	sqlStr := "select user_name from blog_user where uid = ?"
	rows := DB.QueryRow(sqlStr, id)
	if rows.Err() != nil {
		log.Panicln(rows.Err())
	}
	_ = rows.Scan(&name)
	return
}

func GetUser(userName, passwd string) *models.User {
	sqlStr := "select * from blog_user where user_name=? and passwd =? limit 1"
	rows := DB.QueryRow(sqlStr, userName, passwd)
	if rows.Err() != nil {
		log.Println(rows.Err())
		return nil
	}
	var user = &models.User{}
	err := rows.Scan(&user.Uid, &user.UserName, &user.Passwd, &user.Avatar, &user.CreateAt, &user.UpdateAt)
	if err != nil {
		log.Println(rows.Err())
		return nil
	}
	return user
}
