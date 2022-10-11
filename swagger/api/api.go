/*
 * @Author: GG
 * @Date: 2022-10-10 15:38:07
 * @LastEditTime: 2022-10-11 16:29:09
 * @LastEditors: GG
 * @Description:
 * @FilePath: \swagger\api\api.go
 *
 */
package api

import (
	"net/http"
	"strconv"
	"swagger/config"
	"swagger/models"

	"github.com/gin-gonic/gin"
)

// @Summary      查询
// @Description  查询
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param limit body string true "y" minlength(3) maxlength(100)
// @Param offset body string true "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success      200  {object} models.Response
// @Failure      400  {object} models.Response
// @Failure      404  {object} models.Response
// @Failure      500  {object} models.Response
// @Router       /posts [get]
func Posts(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	var posts []models.Post

	config.DB.Limit(limit).Offset(offset).Find(&posts)
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"data":    posts,
	})
}

// @Summary      查询
// @Description  查询
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param id path int true "pid"
// @Success      200  {object} models.Response
// @Failure      400  {object} models.Response
// @Failure      404  {object} models.Response
// @Failure      500  {object} models.Response
// @Router       /posts/{id} [get]
func Show(c *gin.Context) {
	post := getById(c)
	if post.ID == 0 {
		c.JSON(http.StatusOK, models.Response{
			Msg:  "post not found",
			Data: "",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Msg:  "",
		Data: post,
	})
}

// @Summary      添加post
// @Description  添加post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param content body string true "json"
// @Success      200  {object} models.Response
// @Failure      400  {object} models.Response
// @Failure      404  {object} models.Response
// @Failure      500  {object} models.Response
// @Router       /posts [post]
func Store(c *gin.Context) {

	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Msg:  err.Error(),
			Data: "",
		})
		return
	}
	post.Status = "Active"
	res := config.DB.Create(&post)
	if res.Error != nil {
		c.JSON(http.StatusForbidden, models.Response{
			Msg:  res.Error.Error(),
			Data: post,
		})
	}
	c.JSON(http.StatusOK, models.Response{
		Msg:  "",
		Data: post,
	})
}

// @Summary      修改post
// @Description  修改post
// @Tags         posts
// @Accept       html
// @Produce      json
// @Param id path int true "pid"
// @Param content body string true "json"
// @Success      200  {object} models.Response
// @Failure      400  {object} models.Response
// @Failure      404  {object} models.Response
// @Failure      500  {object} models.Response
// @Router       /posts/{id} [patch]
func Update(c *gin.Context) {
	oldpost := getById(c)
	if oldpost.ID == 0 {
		c.JSON(http.StatusOK, models.Response{
			Msg:  "post not found",
			Data: "",
		})
		return
	}
	var newpost models.Post
	if err := c.ShouldBind(&newpost); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Msg:  err.Error(),
			Data: "",
		})
		return
	}

	oldpost.Title = newpost.Title
	oldpost.Des = newpost.Des

	if newpost.Status != "" {
		oldpost.Status = newpost.Status
	}
	config.DB.Save(&oldpost)
	c.JSON(http.StatusOK, models.Response{
		Msg:  "Post has been updated",
		Data: oldpost,
	})
}

// @Summary      删除post
// @Description  删除post
// @Tags         posts
// @Accept       json
// @Produce      json
// @Param id path int true "id"
// @Success      200  {object} models.Response
// @Failure      400  {object} models.Response
// @Failure      404  {object} models.Response
// @Failure      500  {object} models.Response
// @Router       /posts/{id} [delete]
func Delete(c *gin.Context) {
	post := getById(c)
	config.DB.Unscoped().Delete(&post)
	c.JSON(http.StatusOK, models.Response{
		Msg:  "delete successfuly",
		Data: "",
	})
}

func getById(c *gin.Context) models.Post {
	var post models.Post
	id := c.Param("id")
	config.DB.First(&post, id)
	return post
}
