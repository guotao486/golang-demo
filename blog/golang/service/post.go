/*
 * @Author: GG
 * @Date: 2022-08-19 14:34:23
 * @LastEditTime: 2022-08-20 10:44:22
 * @LastEditors: GG
 * @Description:post service
 * @FilePath: \golang-demo\blog\golang\service\post.go
 *
 */
package service

import (
	"fmt"
	"golang-demo/blog/golang/config"
	"golang-demo/blog/golang/dao"
	"golang-demo/blog/golang/models"
	"html/template"
	"log"
)

func GetPostDetail(pId int) (*models.PostRes, error) {
	post, err := dao.GetPostById(pId)
	if err != nil {
		return nil, err
	}
	userName := dao.GetUserNameById(post.UserId)
	categoryName := dao.GetCategoryNameById(post.CategoryId)
	var postMore = models.PostMore{
		post.Pid,
		post.Title,
		post.Slug,
		template.HTML(post.Content),
		post.CategoryId,
		categoryName,
		post.UserId,
		userName,
		post.ViewCount,
		post.Type,
		models.DateDay(post.CreateAt),
		models.DateDay(post.UpdateAt),
	}
	var postRes = &models.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}

	return postRes, nil
}

func Writing() (*models.WritingRes, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var wr = &models.WritingRes{
		Title:     config.Cfg.Viewer.Title,
		CdnURL:    config.Cfg.System.CdnURL,
		Categorys: categorys,
	}
	fmt.Printf("wr.Categorys: %v\n", wr.Categorys)
	return wr, nil
}
