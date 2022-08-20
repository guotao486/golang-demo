/*
 * @Author: GG
 * @Date: 2022-08-20 10:57:15
 * @LastEditTime: 2022-08-20 13:41:35
 * @LastEditors: GG
 * @Description:post api
 * @FilePath: \golang-demo\blog\golang\api\post.go
 *
 */
package api

import (
	"errors"
	"fmt"
	"golang-demo/blog/golang/common"
	"golang-demo/blog/golang/models"
	"golang-demo/blog/golang/service"
	"golang-demo/blog/golang/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*JSONApi) GetPost(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Path
	pIdStr := strings.TrimPrefix(s, "/api/v1/post/")
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		log.Println(err)
		common.Error(w, errors.New("不识别此请求路径"))
		return
	}

	post, err := service.GetPostById(pid)
	if err != nil {
		log.Println(err)
		common.Error(w, err)
		return
	}
	common.Success(w, post)
}

func (*JSONApi) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {

	// 判断用户是否登录
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		log.Println(err)
		common.Error(w, errors.New("登录已失效"))
		return
	}
	uid := claim.Uid

	//post save
	method := r.Method
	params := common.GetRequestJsonParam(r)
	switch method {
	case http.MethodPost:
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pType := int(postType)
		post := &models.Post{
			-1,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		_, err := service.SavePost(post)
		if err != nil {
			log.Println(err)
			common.Error(w, errors.New("保存失败"))
		}
		common.Success(w, post)
	case http.MethodPut:
		fmt.Printf("params: %v\n", params["categoryId"])
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := params["type"].(float64)
		pidFloat := params["pid"].(float64)
		pType := int(postType)
		pId := int(pidFloat)
		post := &models.Post{
			pId,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		_, err := service.UpdatePost(post)
		if err != nil {
			log.Println(err)
			common.Error(w, errors.New("修改失败"))
		}
		common.Success(w, post)
	}
}
