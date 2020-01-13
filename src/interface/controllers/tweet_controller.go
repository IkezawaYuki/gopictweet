package controllers

import (
	"fmt"
	"github.com/IkezawaYuki/gopictweet/src/domain"
	"github.com/IkezawaYuki/gopictweet/src/usecase/interactor"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TweetController struct {
	tweetInteractor interactor.TweetInteractor
}

func (t *TweetController) CreateTweet(c *gin.Context) {
	tweet := domain.Tweet{}
	c.Bind(&tweet)

	// todo session確認

	result, err := t.tweetInteractor.Create(&tweet)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/")
	}
	fmt.Println(result)
	c.Redirect(http.StatusFound, "/")
}
