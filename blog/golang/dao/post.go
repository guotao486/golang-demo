/*
 * @Author: GG
 * @Date: 2022-08-18 11:32:15
 * @LastEditTime: 2022-08-22 11:01:54
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
func CountGetAllPostBySlug(slug string) (count int) {
	sqlStr := "select count(1) from blog_post where slug=?"
	rows := DB.QueryRow(sqlStr, slug)
	_ = rows.Scan(&count)
	return
}
func GetAllPost() ([]models.Post, error) {
	sqlStr := "select * from blog_post"
	rows, err := DB.Query(sqlStr)
	if err != nil {
		log.Println("GetAllPost 查询错误：", err)
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
func GetPostPageBySlug(slug string, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	sqlStr := "select * from blog_post where slug=? limit ?,?"
	rows, err := DB.Query(sqlStr, slug, page, pageSize)
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

func SavePost(post *models.Post) (bool, error) {
	sqlStr := "insert into blog_post (title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) values(?,?,?,?,?,?,?,?,?,?)"
	row, err := DB.Exec(sqlStr, post.Title, post.Content, post.Markdown, post.CategoryId, post.UserId, post.ViewCount, post.Type, post.Slug, post.CreateAt, post.UpdateAt)
	if err != nil {
		log.Println("SavePost 执行错误：", err)
		return false, err
	}
	pid, _ := row.LastInsertId()
	post.Pid = int(pid)
	return true, nil
}

func UpdatePost(post *models.Post) (bool, error) {
	sqlStr := "update blog_post set title=?,content=?,markdown=?,category_id=?,type=?,slug=?,update_at=? where pid=?"
	row, err := DB.Exec(sqlStr, post.Title, post.Content, post.Markdown, post.CategoryId, post.Type, post.Slug, post.UpdateAt, post.Pid)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if count, err := row.RowsAffected(); count < 1 {
		return false, err
	}
	return true, nil
}

func GetPostByTitle(condtion string) ([]models.Post, error) {
	sqlStr := "select pid,title from blog_post where title like ?"
	rows, err := DB.Query(sqlStr, "%"+condtion+"%")
	if err != nil {
		log.Println("GetPostByTitle 查询错误:", err)
		return nil, err
	}

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		rows.Scan(&post.Pid, &post.Title)

		posts = append(posts, post)
	}

	return posts, nil
}
