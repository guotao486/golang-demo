/*
 * @Author: GG
 * @Date: 2022-08-22 14:35:52
 * @LastEditTime: 2022-08-22 16:36:59
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gorm\models\users.go
 *
 */
package models

import (
	"database/sql"
	"time"
)

// 自定义配置，声明模型
type User struct {
	ID           uint
	Name         string
	Email        string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreateAt     time.Time
	UpdateAt     time.Time
	DeleteAt     time.Time
}
