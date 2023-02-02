package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

/**
 * @description: 文件hash
 * @param {string} filePath
 * @return {*}
 */
func FileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

/**
 * @description: 字符串hash
 * @param {string} value
 * @return {*}
 */
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

func main() {

}
