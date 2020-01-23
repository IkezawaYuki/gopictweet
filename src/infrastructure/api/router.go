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
	router.Static("/assets", "./public")
	router.LoadHTMLGlob("src/infrastructure/view/*.html")

	router.HTMLRender = createMyRender()

	handler, err := gorm.Open("mysql", "root:@/pictweet?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	controller := controllers.NewPicTweetController(handler)

	router.GET("/tweet/edit", controller.EditTweet)
	router.GET("/", controller.Index)

	router.Run(":8081")
}

func createMyRender() multitemplate.Render {
	r := multitemplate.New()
	path := "src/infrastructure/view/"
	//
	//r.AddFromFiles("editTweet", path+"layout.html", path+"private.navbar.html", path+"edit.tweet.html")
	//r.AddFromFiles("Index", path+"layout.html", path+"public.navbar.html", path+"index.html")
	//r.AddFromFiles("Index_Private", path+"layout.html", path+"private.navbar.html", path+"index.html")
	r.AddFromFiles("article", path+"base.html", path+"index.html", path+"article.html")

	return r
}
