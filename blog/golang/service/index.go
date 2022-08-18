/*
 * @Author: GG
 * @Date: 2022-08-18 10:40:23
 * @LastEditTime: 2022-08-18 15:52:26
 * @LastEditors: GG
 * @Description:category service
 * @FilePath: \golang-demo\blog\golang\service\index.go
 *
 */
package service

import (
	"golang-demo/blog/golang/config"
	"golang-demo/blog/golang/dao"
	"golang-demo/blog/golang/models"
	"html/template"
)

func GetAllIndexInfo(page, pageSize int) (*models.HomeResponse, error) {
	categorys, err := dao.GelAllCategory()
	if err != nil {
		return nil, err
	}

	// 获取文章部分
	posts, err := dao.GetPostPage(page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
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
	total := dao.CountGetAllPost()
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
	return hr, nil
}
