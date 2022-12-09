/*
 * @Author: GG
 * @Date: 2022-12-09 14:40:24
 * @LastEditTime: 2022-12-09 15:08:04
 * @LastEditors: GG
 * @Description: 模板方法模式
 * @FilePath: \设计模式\模板方法模式\main.go
 *
 */
package main

import "fmt"

// 模板方法模式定义了一个算法的步骤，并允许子类，为一个或多个步骤提供其实践方式。让子类别在不改变算法架构的情况下，重新定义算法中的某些步骤
// 利用了继承特性

// 下载接口
type Downloader interface {
	Download(uri string)
}

// 具体实施接口
type implement interface {
	download()
	save()
}

// 模板
type template struct {
	implement
	uri string
}

// 实例化模板
func NewTemplate(impl implement) *template {
	return &template{
		implement: impl,
		uri:       "",
	}
}

// 实现下载接口 Download方法
func (t template) Download(uri string) {
	t.uri = uri
	fmt.Println("准备下载")
	t.implement.download()
	t.implement.save()
	fmt.Println("下载完毕")
}

func (t template) save() {
	fmt.Println("默认保存")
}

// ===快播下载器===
type QuickDownloader struct {
	*template
}

// 实例化快播下载器
func NewQuickDownloader() *QuickDownloader {
	q := &QuickDownloader{}
	t := NewTemplate(q)
	q.template = t
	return q
}

func (q *QuickDownloader) download() {
	fmt.Println("正在使用快播下载器：", q.uri)
}

func (q *QuickDownloader) save() {
	fmt.Println("快播保存")
}

// ==== ftp下载 ====
type FtpDownloader struct {
	*template
}

// 实例化ftp下载器
func NewFtpDownloader() *FtpDownloader {
	f := &FtpDownloader{}
	t := NewTemplate(f)
	f.template = t
	return f
}

func (f *FtpDownloader) download() {
	fmt.Println("正在使用ftp下载器", f.uri)
}

func (f *FtpDownloader) save() {
	fmt.Println("ftp保存")
}

func main() {
	qd := NewQuickDownloader()
	qd.Download("www.baidu.com/zip")

	fp := NewFtpDownloader()
	fp.Download("www.baidu.com/rar")
}
