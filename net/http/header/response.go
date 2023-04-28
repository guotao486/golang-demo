package main

import (
	"fmt"
	"net/http"
	"os"
)

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	// 获取文件名
	filename := "example.pdf"

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err.Error())
		return
	}
	defer file.Close()

	// 设置响应头
	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)

	// 写入响应体
	_, err = file.WriteTo(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %s", err.Error())
	}
}

func main() {
	http.HandleFunc("/download", downloadHandler)
	http.ListenAndServe(":8080", nil)
}
