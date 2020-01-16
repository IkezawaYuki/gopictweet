package controllers

import (
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
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	} else {
		user, err := t.picInteractor.FindUserBySession(ses)
		if err != nil {
			panic(err)
		}
		text := c.PostForm("text")
		image := c.PostForm("image")
		if _, err := t.tweetInteractor.Create(user.Id, text, image); err != nil {
			panic(err)
		}
		c.Redirect(http.StatusCreated, "/")
		return
	}
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
