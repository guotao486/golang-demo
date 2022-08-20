/*
 * @Author: GG
 * @Date: 2022-08-18 10:31:58
 * @LastEditTime: 2022-08-20 09:59:51
 * @LastEditors: GG
 * @Description:category dao
 * @FilePath: \golang-demo\blog\golang\dao\category.go
 *
 */
package dao

import (
	"golang-demo/blog/golang/models"
	"log"
)

func GetAllCategory() ([]models.Category, error) {
	sqlStr := "select * from blog_category"
	rows, err := DB.Query(sqlStr)
	if err != nil {
		log.Panicln("GetAllCategory 查询出错:", err)
		return nil, err
	}

	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Panicln("GetAllCategory 取值出错:", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}

func GetCategoryNameById(id int) (name string) {
	sqlStr := "select name from blog_category where cid = ?"
	rows := DB.QueryRow(sqlStr, id)
	if rows.Err() != nil {
		log.Panicln(rows.Err())
	}
	_ = rows.Scan(&name)
	return
}
