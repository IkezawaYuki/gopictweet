package api

import (
	"github.com/IkezawaYuki/gopictweet/src/interface/controllers"
	"github.com/gin-gonic/contrib/renders/multitemplate"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Run() {

	router.LoadHTMLGlob("src/infrastructure/view")

	router.GET("/ping", controllers.Ping)

	templates := multitemplate.New()
	templates.AddFromFiles("editTweet", "layout.html", "private.navbar.html", "edit.tweet.html")

	router.Run(":8081")
}
