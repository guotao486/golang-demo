/*
 * @Author: GG
 * @Date: 2023-05-04 14:57:11
 * @LastEditTime: 2023-05-04 15:15:56
 * @LastEditors: GG
 * @Description:
 * @FilePath: \xlsx\dao\shopyy.go
 *
 */
package dao

import (
	"errors"

	"gorm.io/gorm"
)

type Product struct {
	ID            uint   `gorm:"id"`
	Title         string `gorm:"title"`
	Price         string `gorm:"price"`
	OriginalPrice string `gorm:"originalPrice"`
	Description   string `gorm:"description"`
	Attribute     string `gorm:"attribute"`
	Images        string `gorm:"images"`
}

func (p Product) TableName() string {
	return "shopyy_product"
}

type ProductDao struct {
	db *gorm.DB
}

func NewProductDao(db *gorm.DB) *ProductDao {
	return &ProductDao{db: db}
}

func (p ProductDao) GetAll() ([]*Product, error) {
	var products []*Product
	result := p.db.Find(&products)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return products, nil
}
