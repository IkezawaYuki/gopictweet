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

// CreateTweet ツイートの作成
func (t *TweetController) CreateTweet(c *gin.Context) {
	ses, err := t.session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
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

// UpdateTweet ツイートの更新
func (t *TweetController) UpdateTweet(c *gin.Context) {
	ses, err := t.session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	user, err := t.picInteractor.FindUserBySession(ses)
	if err != nil {
		panic(err)
	}
	uuid := c.PostForm("uuid")
	text := c.PostForm("text")
	image := c.PostForm("image")
	if _, err := t.tweetInteractor.Update(user.Id, uuid, text, image); err != nil {
		panic(err)
	}
	c.Redirect(http.StatusCreated, "/")
	return

}

func (t *TweetController) NewTweet(c *gin.Context) {
	_, err := t.session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
	} else {
		// htmlの生成

	}
}

func (t *TweetController) EditTweet(c *gin.Context) {
	uuid := c.Query("id")
	tweet, err := t.tweetInteractor.FindByUUID(uuid)
	if err != nil {

		// todo メッセージ：tweetが見つかりません。
	}
	_, err = t.session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
	}
	fmt.Println(tweet)

}

func (t *TweetController) session(c *gin.Context) (ses *domain.Session, err error) {
	cookie, err := c.Cookie("_cookie")
	if err == nil {
		ses, err = t.picInteractor.CheckSession(cookie)
	}
	return nil, err
}
