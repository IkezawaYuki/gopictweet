package api

import (
	"github.com/IkezawaYuki/pictweet-api/src/interface/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var (
	router = gin.Default()
)

func Run() {
	handler, err := gorm.Open("mysql", "root:@/pictweet?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	controller := controllers.NewPicTweetController(handler)

	router.POST("/tweet/create", controller.CreateTweet)
	router.GET("/tweet/edit", controller.EditTweet)
	router.GET("/", controller.Index)

	router.Run(":8081")
}
