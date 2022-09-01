/*
 * @Author: GG
 * @Date: 2022-09-01 17:23:02
 * @LastEditTime: 2022-09-01 17:37:53
 * @LastEditors: GG
 * @Description:
 * @FilePath: \golang-demo\gin\demo\upload\main.go
 *
 */
package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func DoUpload(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func Upload(c *gin.Context) {
	file, _ := c.FormFile("file")
	fmt.Printf("file.Filename: %v\n", file.Filename)
	fmt.Printf("file.Header: %v\n", file.Header)
	fmt.Printf("file.Size: %v\n", file.Size)

	// 上传文件到项目根目录，使用原文件名
	c.SaveUploadedFile(file, "./gin/demo/upload/"+file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
func main() {
	e := gin.Default()
	viewPath, _ := os.Getwd()
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	e.MaxMultipartMemory = 8 << 20 // 8 MiB
	e.LoadHTMLGlob(viewPath + "/gin/demo/templates/uploads/*")
	e.GET("/upload", DoUpload)
	e.POST("/upload", Upload)
	e.Run()
}
