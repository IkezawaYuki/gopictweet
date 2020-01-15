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
	picInteractor   interactor.PictweetInteractor
}

func (t *TweetController) CreateTweet(c *gin.Context) {
	tweet := domain.Tweet{}
	c.Bind(&tweet)
	ses, err := t.session(c)

	// todo tweetオブジェクトを作成

	result, err := t.tweetInteractor.Create(&tweet)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusFound, "/")
	}
	fmt.Println(result)
	c.Redirect(http.StatusFound, "/")
}

func (t *TweetController) UpdateTweet(c *gin.Context) {

}

func (t *TweetController) session(c *gin.Context) (ses *domain.Session, err error) {
	cookie, err := c.Cookie("_cookie")
	if err == nil {
		ses, err = t.picInteractor.CheckSession(cookie)
	}
	return nil, err
}
