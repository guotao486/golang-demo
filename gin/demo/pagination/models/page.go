/*
 * @Author: GG
 * @Date: 2022-09-23 16:09:38
 * @LastEditTime: 2022-09-24 11:14:44
 * @LastEditors: GG
 * @Description:
 * @FilePath: \pagination\models\page.go
 *
 */
package models

type Pagination struct {
	Limit     int    `json:"limit"`     // 每页数量
	Page      int    `json:"page"`      // 当前页码
	Sort      string `json:"sort"`      // 排序字段
	Total     int    `json:"total"`     // 总数量
	TotalPage int    `json:"totalPage"` // 总页数
	LastPage  bool   `json:"lastPage"`  // 最后一页
}

// 计算总页数
func (p *Pagination) GetTotalPage() int {
	if p.Total == 0 {
		return 0
	}
	p.TotalPage = (p.Total + p.Limit - 1) / p.Limit
	return p.TotalPage
}

// 是否最后一页
func (p *Pagination) IsLastPage() bool {
	p.LastPage = false
	if p.TotalPage == 0 || p.TotalPage == p.Page {
		p.LastPage = true
		return p.LastPage
	}

	return p.LastPage
}
