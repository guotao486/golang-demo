/*
 * @Author: GG
 * @Date: 2022-08-19 11:08:56
 * @LastEditTime: 2022-08-20 10:01:05
 * @LastEditors: GG
 * @Description:categoty service
 * @FilePath: \golang-demo\blog\golang\service\category.go
 *
 */
package service

import (
	"golang-demo/blog/golang/config"
	"golang-demo/blog/golang/dao"
	"golang-demo/blog/golang/models"
	"html/template"
)

func GetPostByCategoryId(categoryId, page, pageSize int) (*models.CategoryResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}

	// 获取文章部分
	posts, err := dao.GetPostPageByCateegoryId(categoryId, page, pageSize)

	var postMores []models.PostMore
	categoryName := dao.GetCategoryNameById(categoryId)
	for _, post := range posts {
		userName := dao.GetUserNameById(post.UserId)
		// 最多显示100个字符
		s := []rune(post.Content)
		if len(s) > 100 {
			s = s[0:100]
		}

		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(s),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}

		postMores = append(postMores, postMore)
	}

	// 总数量
	total := dao.CountGetAllPostByCategoryId(categoryId)
	// 页数
	pageCount := (total-1)/pageSize + 1
	// 页码切片
	var pages []int
	for i := 0; i < pageCount; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page != pageCount,
	}
	categoryReesponse := &models.CategoryResponse{
		hr,
		categoryName,
	}
	return categoryReesponse, nil
}
