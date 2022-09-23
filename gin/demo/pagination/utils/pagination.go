/*
 * @Author: GG
 * @Date: 2022-09-23 16:13:29
 * @LastEditTime: 2022-09-23 17:40:41
 * @LastEditors: GG
 * @Description:
 * @FilePath: \pagination\utils\pagination.go
 *
 */
package utils

import (
	"fmt"
	"pagination/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// url?limit=2&page=2
func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	limit := 2
	page := 1
	sort := "created_at asc"

	// url?limit=2&page=2
	query := c.Request.URL.Query()
	for key, value := range query {
		fmt.Printf("value: %v\n", value)

		// value: [20]
		// value: [30]
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}
	}

	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}

// url/:limit/:page
func GeneratePaginationFromRequestParams(c *gin.Context) models.Pagination {
	limit := 2
	page := 1
	sort := "created_at asc"

	// [{limit 1} {page}]
	query := c.Params
	// query: map[limit:[20] page:[30]]
	for _, value := range query {
		// value: {limit 1}
		// value: {page 1}
		fmt.Printf("value: %v\n", value)

		switch value.Key {
		case "limit":
			limit, _ = strconv.Atoi(value.Value)
			break
		case "page":
			page, _ = strconv.Atoi(value.Value)
			break
		case "sort":
			sort = value.Value
			break
		}
	}

	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
