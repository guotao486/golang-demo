package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

var dir, _ = os.Getwd()
var uploadPath = path.Join(dir, "uploads")
var uploadTempPath = path.Join(uploadPath, "temp")

// 加载html前段页面
func home(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, "")
	return
}

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func uploadFile(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	index := r.PostFormValue("index")
	hash := r.PostFormValue("hash")
	// 获取uploads下所有的文件夹
	nameList, err := ioutil.ReadDir(uploadPath)
	m := map[string]interface{}{
		"code": 46900,
		"msg":  "文件已上传",
	}
	result, _ := json.MarshalIndent(m, "", "    ")
	// 循环判断hash是否在文件里如果有就返回上传已完成
	for _, name := range nameList {
		tmpName := strings.Split(name.Name(), "_")[0]
		if tmpName == hash {
			fmt.Fprintf(w, string(result))
			return
		}
	}

	// 块文件地址
	chunksPath := path.Join(uploadTempPath, hash, "/")

	// 判断块文件目录是否存在，不在则创建
	isPathExists, err := PathExists(chunksPath)
	if !isPathExists {
		err = os.MkdirAll(chunksPath, os.ModePerm)
	}
	// 打开文件写入，没则创建
	destFile, err := os.OpenFile(path.Join(chunksPath+"/"+hash+"-"+index), syscall.O_CREAT|syscall.O_WRONLY, 0777)
	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(destFile)
	buf := make([]byte, 1024*1024) // 1M buf
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			writer.Flush()
			break
		} else if err != nil {
			return
		} else {
			writer.Write(buf[:n])
		}
	}

	defer file.Close()
	defer destFile.Close()
	if err != nil {
		log.Fatal("%v", err)
	}
	fmt.Printf("第%s:%s块上传完成\n", index, destFile.Name())
}

// 合并文件
func chunks(w http.ResponseWriter, r *http.Request) {
	size, _ := strconv.ParseInt(r.PostFormValue("size"), 10, 64)
	hash := r.PostFormValue("hash")
	name := r.PostFormValue("name")

	toSize, _ := getDirSize(path.Join(uploadTempPath, hash, "/"))
	if size != toSize {
		fmt.Fprintf(w, "文件上传错误")
	}
	chunksPath := path.Join(uploadTempPath, hash, "/")
	files, _ := ioutil.ReadDir(chunksPath)
	// 排序
	filesSort := make(map[string]string)
	for _, f := range files {
		nameArr := strings.Split(f.Name(), "-")
		filesSort[nameArr[1]] = f.Name()
	}
	saveFile := path.Join(uploadPath, name)
	if exists, _ := PathExists(saveFile); exists {
		os.Remove(saveFile)
	}
	fs, _ := os.OpenFile(saveFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	var wg sync.WaitGroup
	filesCount := len(files)
	if filesCount != len(filesSort) {
		fmt.Fprintf(w, "文件上传错误2")
	}
	wg.Add(filesCount)
	for i := 0; i < filesCount; i++ {
		// 这里一定要注意按顺序读取不然文件就会损坏
		fileName := path.Join(chunksPath, "/"+filesSort[strconv.Itoa(i)])
		data, err := ioutil.ReadFile(fileName)
		fmt.Println(err)
		fs.Write(data)

		wg.Done()
	}
	wg.Wait()
	os.RemoveAll(path.Join(chunksPath, "/"))
	m := map[string]interface{}{
		"code": 20000,
		"msg":  "上传成功",
	}
	result, _ := json.MarshalIndent(m, "", "    ")
	fmt.Fprintf(w, string(result))
	defer fs.Close()

}

// 获取整体文件夹大小
func getDirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
func main() {
	http.HandleFunc("/", home) // set router
	http.HandleFunc("/uploadFile", uploadFile)
	http.HandleFunc("/file/chunks", chunks)
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("Error while starting GO http server on port - 8080 : ", err) //log error and exit in case of error at server boot up
	}
}
