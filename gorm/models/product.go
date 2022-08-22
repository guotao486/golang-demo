/*
 * @Author: GG
 * @Date: 2022-08-22 14:38:31
 * @LastEditTime: 2022-08-22 14:52:23
 * @LastEditors: GG
 * @Description:gorm product models
 * @FilePath: \golang-demo\gorm\models\product.go
 *
 */
package models

import "gorm.io/gorm"

/*
gorm.Model
ID        uint `gorm:"primarykey"`
CreatedAt time.Time
UpdatedAt time.Time
DeletedAt DeletedAt `gorm:"index"`
*/
// 使用gorm约定model
type Product struct {
	gorm.Model
	Code  string
	Price uint
}
