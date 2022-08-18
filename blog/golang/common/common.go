/*
 * @Author: GG
 * @Date: 2022-08-18 16:18:08
 * @LastEditTime: 2022-08-18 18:12:31
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\blog\golang\common\common.go
 *
 */
package common

import (
	"encoding/json"
	"golang-demo/blog/golang/models"
	"io/ioutil"
	"log"
	"net/http"
)

// 获取http.request请求的json参数
func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}

	// 读取body内容
	body, _ := ioutil.ReadAll(r.Body)
	// 将读取的body转换map
	_ = json.Unmarshal(body, &params)
	return params
}

func Error(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = 400
	result.Error = err.Error()
	result.Data = ""
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}
