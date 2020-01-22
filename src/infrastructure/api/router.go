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
	path := "src/infrastructure/view/"
	templates.AddFromFiles("editTweet", path+"layout.html", path+"private.navbar.html", path+"edit.tweet.html")
	templates.AddFromFiles("Index", path+"layout.html", path+"public.navbar.html", path+"index.html")
	templates.AddFromFiles("Index_Private", path+"layout.html", path+"private.navbar.html", path+"index.html")

	handler, err := gorm.Open("mysql", "root:@/pictweet?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	controller := controllers.NewPicTweetController(handler)

	router.GET("/tweet/edit", controller.EditTweet)
	router.GET("/", controller.Index)

	router.Run(":8081")
}
