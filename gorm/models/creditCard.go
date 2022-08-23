/*
 * @Author: GG
 * @Date: 2022-08-23 15:31:06
 * @LastEditTime: 2022-08-23 15:40:04
 * @LastEditors: GG
 * @Description:creditCard models
 * @FilePath: \golang-demo\gorm\models\creditCard.go
 *
 */
package models

import (
	"fmt"

	"gorm.io/gorm"
)

type CreditCard struct {
	gorm.Model
	Number string
	UserId uint
}

func (c *CreditCard) AfterCreate(tx *gorm.DB) (err error) {

	fmt.Println("CreditCard Models AfterCreate")
	fmt.Printf("c.ID: %v\n", c.ID)
	fmt.Printf("c.UserId: %v\n", c.UserId)
	return
}
