package main

/**特别注意，生成docs包后需要导入**/
/**特别注意，生成docs包后需要导入**/
/**特别注意，生成docs包后需要导入**/
/**特别注意，生成docs包后需要导入**/
/**特别注意，生成docs包后需要导入**/
/**特别注意，生成docs包后需要导入**/
/**特别注意，生成docs包后需要导入**/
import (
	"swagger/api"
	"swagger/config"
	_ "swagger/docs"
	"swagger/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var err error

// @title gin+gorm crud 测试swagger（必填）
// @version 1.0 （必填）
// @description gin+gorm crud 测试swagger
// @license.name Apache 2.0
// @contact.name go-swagger帮助文档
// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @host localhost:8000
// @BasePath /
func main() {
	//establishing connection with mysql database 'CRUD'
	dsn := "root:123456@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	//handle the error comes from the connection with DB
	if err != nil {
		panic(err.Error())
	}

	config.DB = db
	db.AutoMigrate(&models.Post{})

	server := gin.Default()

	//set up the different routes
	server.GET("/posts", api.Posts)
	server.GET("/posts/:id", api.Show)
	server.POST("/posts", api.Store)
	server.PATCH("/posts/:id", api.Update)
	server.DELETE("/posts/:id", api.Delete)

	server.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//start the server and listen on the port 8000
	server.Run(":8000")
}
