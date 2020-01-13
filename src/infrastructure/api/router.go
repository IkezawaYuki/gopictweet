package api

import (
	"github.com/IkezawaYuki/gopictweet/src/interface/controllers"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Run() {
	router.GET("/ping", controllers.Ping)

	router.Run(":8081")
}
