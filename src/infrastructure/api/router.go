package api

import (
	"github.com/IkezawaYuki/gopictweet/src/interface/controllers"
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var (
	router = gin.Default()
)

func Run() {

	router.LoadHTMLGlob("src/infrastructure/view/*.html")

	templates := multitemplate.New()
	templates.AddFromFiles("editTweet", "src/infrastructure/view/layout.html", "src/infrastructure/view/private.navbar.html", "src/infrastructure/view/edit.tweet.html")

	handler, err := gorm.Open("mysql", "root:@/pictweet?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	controller := controllers.NewPicTweetController(handler)

	router.GET("/tweet/edit", controller.EditTweet)

	router.Run(":8081")
}
