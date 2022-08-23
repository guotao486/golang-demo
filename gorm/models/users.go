/*
 * @Author: GG
 * @Date: 2022-08-22 14:35:52
 * @LastEditTime: 2022-08-23 16:05:21
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gorm\models\users.go
 *
 */
package models

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// 自定义配置，声明模型
type User struct {
	ID           uint
	Name         string `gorm:"default:golang"`
	Email        string `gorm:"default:''"`
	Age          uint8  `gorm:"default:18"`
	Birthday     time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreateAt     time.Time `gorm:"autoCreateTime"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	CreditCard   CreditCard
}

// hook
// 跳过hook DB.Session(&gorm.Session{SkipHooks: true}).Create(&user)
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("User Models Before Create")
	fmt.Printf("u.ID: %v\n", u.ID)
	fmt.Println("----------------------------------")

	return
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("User Models After Create")
	fmt.Printf("u.ID: %v\n", u.ID)
	fmt.Println("----------------------------------")
	return
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	fmt.Println("User Models Before Save")
	fmt.Printf("u.ID: %v\n", u.ID)
	fmt.Println("----------------------------------")
	return
}

func (u *User) AfterSave(tx *gorm.DB) (err error) {
	fmt.Println("User Models After Save")
	fmt.Printf("u.ID: %v\n", u.ID)
	fmt.Println("----------------------------------")
	return
}
