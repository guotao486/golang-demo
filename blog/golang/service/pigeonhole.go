/*
 * @Author: GG
 * @Date: 2022-08-20 13:54:15
 * @LastEditTime: 2022-08-20 14:15:44
 * @LastEditors: GG
 * @Description:pigeonhole service
 * @FilePath: \golang-demo\blog\golang\service\pigeonhole.go
 *
 */
package service

import (
	"fmt"
	"golang-demo/blog/golang/config"
	"golang-demo/blog/golang/dao"
	"golang-demo/blog/golang/models"
)

func FindPostPigeonhole() (*models.PigeonholeRes, error) {
	// 查询所有文章 进行月份整理
	posts, err := dao.GetAllPost()
	if err != nil {
		return nil, err
	}
	pigeonholeMap := make(map[string][]models.Post)
	for _, post := range posts {
		fmt.Printf("post: %v\n", post)
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}
	// 查询所有分类
	categorys, _ := dao.GetAllCategory()

	return &models.PigeonholeRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categorys:    categorys,
		Lines:        pigeonholeMap,
	}, nil
}
