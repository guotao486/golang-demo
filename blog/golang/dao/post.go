/*
 * @Author: GG
 * @Date: 2022-08-18 11:32:15
 * @LastEditTime: 2022-08-19 15:40:25
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\blog\golang\dao\post.go
 *
 */
package dao

import (
	"golang-demo/blog/golang/models"
	"log"
)

func CountGetAllPostByCategoryId(cateogryId int) (count int) {
	sqlStr := "select count(1) from blog_post where category_id = ?"
	row := DB.QueryRow(sqlStr, cateogryId)
	if row.Err() != nil {
		log.Println("CountGetAllPostByCategoryId 查询错误：", row.Err())
	}
	_ = row.Scan(&count)
	return
}

func CountGetAllPost() (count int) {
	sqlStr := "select count(1) from blog_post"
	rows := DB.QueryRow(sqlStr)
	_ = rows.Scan(&count)
	return
}

func GetPostPage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	sqlStr := "select * from blog_post limit ?,?"
	rows, err := DB.Query(sqlStr, page, pageSize)
	if err != nil {
		log.Panicln("GetPostPage 查询错误:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostPageByCateegoryId(categoryId, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	sqlStr := "select * from blog_post where category_id = ? limit ?,?"
	rows, err := DB.Query(sqlStr, categoryId, page, pageSize)
	if err != nil {
		log.Panicln("GetPostPage 查询错误:", err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetPostById(pId int) (*models.Post, error) {
	sqlStr := "select * from blog_post where pid = ?"
	row := DB.QueryRow(sqlStr, pId)
	if row.Err() != nil {
		log.Println("GetPostDetailById 查询失败：", row.Err())
		return nil, row.Err()
	}
	var post = &models.Post{}
	_ = row.Scan(
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt,
	)

	return post, nil
}
