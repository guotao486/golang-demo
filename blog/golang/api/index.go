/*
 * @Author: GG
 * @Date: 2022-08-18 09:59:10
 * @LastEditTime: 2022-08-18 17:38:39
 * @LastEditors: GG
 * @Description:返回json 的index api
 * @FilePath: \golang-demo\blog\golang\api\index.go
 *
 */
package api

import (
	"encoding/json"
	"golang-demo/blog/golang/common"
	"golang-demo/blog/golang/service"
	"net/http"
)

// index 结构体
// `json:title` 表示返回json格式时对应的key
type IndexData struct {
	Title string `json:title`
	Desc  string `json:desc`
}

// 响应对应路径的函数
// http.ResponseWriter 响应对象，http.Request 请求对象
func (*JSONApi) Index(w http.ResponseWriter, r *http.Request) {
	// 设置响应抬头
	w.Header().Set("Content-type", "application/json")

	var indexData IndexData
	indexData.Title = "bolang blog title"
	indexData.Desc = "bolang blog desc"
	jsonStr, _ := json.Marshal(indexData)
	// 响应输出
	// w.Write([]byte("hello world!"))
	w.Write(jsonStr)
}

func (*JSONApi) Login(w http.ResponseWriter, r *http.Request) {
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	passwd := params["passwd"].(string)

	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}

	common.Success(w, loginRes)

}
