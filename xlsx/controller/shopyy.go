package controller

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
	"xlsx/dao"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

type ProductHandler struct {
	dao dao.ProductDao
}

func NewProductHandler(dao dao.ProductDao) *ProductHandler {
	return &ProductHandler{dao: dao}
}

func (h *ProductHandler) Import(c *gin.Context) {
	// 打开文件
	file, err := os.Open("products_import_template.xlsx")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// 读取文件
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	xlFile, err := xlsx.OpenBinary(fileBytes)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {

			if i == 338 {
				fmt.Printf("row.Cells[30].Value: %v\n", row.Cells[30].Value)
				return
			}
		}
	}

}

func (h *ProductHandler) Export(c *gin.Context) {
	products, err := h.dao.GetAll()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// 创建xlsx文件
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("product")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// 表头
	headerRow := sheet.AddRow()
	headerRow.AddCell().Value = "商品ID"
	headerRow.AddCell().Value = "创建时间"
	headerRow.AddCell().Value = "商品标题*"
	headerRow.AddCell().Value = "商品属性*"
	headerRow.AddCell().Value = "商品类型"
	headerRow.AddCell().Value = "商品副标题"
	headerRow.AddCell().Value = "商品描述"
	headerRow.AddCell().Value = "简短描述"
	headerRow.AddCell().Value = "SEO标题"
	headerRow.AddCell().Value = "SEO描述"
	headerRow.AddCell().Value = "SEO URL Handle"
	headerRow.AddCell().Value = "SEO关键词"
	headerRow.AddCell().Value = "商品上架"
	headerRow.AddCell().Value = "商品收税"
	headerRow.AddCell().Value = "商品spu"
	headerRow.AddCell().Value = "虚拟销量值"
	headerRow.AddCell().Value = "跟踪库存"
	headerRow.AddCell().Value = "库存规则*"
	headerRow.AddCell().Value = "专辑名称"
	headerRow.AddCell().Value = "标签"
	headerRow.AddCell().Value = "供应商名称"
	headerRow.AddCell().Value = "款式1"
	headerRow.AddCell().Value = "款式2"
	headerRow.AddCell().Value = "款式3"
	headerRow.AddCell().Value = "商品售价*"
	headerRow.AddCell().Value = "商品原价"
	headerRow.AddCell().Value = "商品SKU"
	headerRow.AddCell().Value = "商品重量(kg)"
	headerRow.AddCell().Value = "商品条形码"
	headerRow.AddCell().Value = "商品库存"
	headerRow.AddCell().Value = "商品图片*"

	spu := 30504
	for _, product := range products {
		row := sheet.AddRow()
		row.AddCell().SetValue("")                  // "商品ID"
		row.AddCell().SetValue("")                  // "创建时间"
		row.AddCell().SetValue(product.Title)       //"商品标题*"
		row.AddCell().SetValue("M")                 //"商品属性*"
		row.AddCell().SetValue("")                  //"商品类型"
		row.AddCell().SetValue("")                  //"商品副标题"
		row.AddCell().SetValue(product.Description) //"商品描述"
		row.AddCell().SetValue("")                  //"简短描述"
		row.AddCell().SetValue("")                  //"SEO标题"
		row.AddCell().SetValue("")                  //"SEO描述"
		row.AddCell().SetValue("")                  //"SEO URL Handle"
		row.AddCell().SetValue("")                  //"SEO关键词"
		row.AddCell().SetValue("Y")                 //"商品上架"
		row.AddCell().SetValue("Y")                 //"商品收税"
		row.AddCell().SetValue(spu)                 //"商品spu"
		row.AddCell().SetValue(0)                   //"虚拟销量值"
		row.AddCell().SetValue("N")                 //"跟踪库存"
		row.AddCell().SetValue(1)                   //"库存规则*"
		row.AddCell().SetValue("")                  // "专辑名称"
		row.AddCell().SetValue("")                  //"标签"
		row.AddCell().SetValue("")                  //"供应商名称"
		row.AddCell().SetValue("Size")              //"款式1"
		row.AddCell().SetValue("")                  //"款式2"
		row.AddCell().SetValue("")                  //"款式3"
		row.AddCell().SetValue(product.Price)       //"商品售价*"
		if len(product.OriginalPrice) == 0 {
			row.AddCell().SetValue(0) //"商品原价"
		} else {
			row.AddCell().SetValue(product.OriginalPrice) //"商品原价"
		}
		row.AddCell().SetValue(fmt.Sprintf("%s-%d", product.Title, spu)) //"商品SKU"
		row.AddCell().SetValue("")                                       //"商品重量(kg)"
		row.AddCell().SetValue("")                                       //"商品条形码"
		row.AddCell().SetValue("")                                       //"商品库存"
		images := strings.Replace(product.Images, "\n", "", 1)
		images = strings.Replace(images, "\r", "", -1)
		images = strings.Replace(images, "\n", ",", -1)

		row.AddCell().SetValue(images) //"商品图片*"
		spu++

		attribute := product.Attribute
		if attribute != "" {
			attributeSlice := strings.Split(attribute, "\n")
			for _, attr := range attributeSlice {
				prow := sheet.AddRow()
				prow.AddCell().SetValue("")            // "商品ID"
				prow.AddCell().SetValue("")            // "创建时间"
				prow.AddCell().SetValue(product.Title) //"商品标题*"
				prow.AddCell().SetValue("P")           //"商品属性*"
				prow.AddCell().SetValue("")            //"商品类型"
				prow.AddCell().SetValue("")            //"商品副标题"
				prow.AddCell().SetValue("")            //"商品描述"
				prow.AddCell().SetValue("")            //"简短描述"
				prow.AddCell().SetValue("")            //"SEO标题"
				prow.AddCell().SetValue("")            //"SEO描述"
				prow.AddCell().SetValue("")            //"SEO URL Handle"
				prow.AddCell().SetValue("")            //"SEO关键词"
				prow.AddCell().SetValue("Y")           //"商品上架"
				prow.AddCell().SetValue("Y")           //"商品收税"
				prow.AddCell().SetValue(spu)           //"商品spu"
				prow.AddCell().SetValue(0)             //"虚拟销量值"
				prow.AddCell().SetValue("N")           //"跟踪库存"
				prow.AddCell().SetValue(1)             //"库存规则*"
				prow.AddCell().SetValue("")            // "专辑名称"
				prow.AddCell().SetValue("")            //"标签"
				prow.AddCell().SetValue("")            //"供应商名称"
				prow.AddCell().SetValue(attr)          //"款式1"
				prow.AddCell().SetValue("")            //"款式2"
				prow.AddCell().SetValue("")            //"款式3"
				prow.AddCell().SetValue(product.Price) //"商品售价*"
				if len(product.OriginalPrice) == 0 {
					prow.AddCell().SetValue(0) //"商品原价"
				} else {
					prow.AddCell().SetValue(product.OriginalPrice) //"商品原价"
				}
				prow.AddCell().SetValue(fmt.Sprintf("%s-%d", product.Title, spu)) //"商品SKU"
				prow.AddCell().SetValue("")                                       //"商品重量(kg)"
				prow.AddCell().SetValue("")                                       //"商品条形码"
				prow.AddCell().SetValue("")                                       //"商品库存"
				prow.AddCell().SetValue("")                                       //"商品图片*"
				spu++
			}
		}
	}
	// 文件保存到临时目录
	fileName := fmt.Sprintf("shopyy-product-%d.xlsx", time.Now().Unix())
	filePath := filepath.Join(os.TempDir(), fileName)
	if err := file.Save(filePath); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// Excel 文件下载
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.File(filePath)
}
