package controllers

import (
	"fmt"
	"github.com/IkezawaYuki/gopictweet/src/domain/model"
	"github.com/IkezawaYuki/gopictweet/src/interface/adapter"
	"github.com/IkezawaYuki/gopictweet/src/usecase/inputport"
	"github.com/IkezawaYuki/gopictweet/src/usecase/interactor"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type PictweetController struct {
	tweetInteractor   inputport.TweetInputport
	userInteractor    inputport.UserInputport
	commentInteractor inputport.CommentInputport
	// todo プレゼンターの追加が必要か。
}

func NewPicTweetController(handler *gorm.DB) *PictweetController {
	return &PictweetController{
		tweetInteractor: interactor.NewTweetInteractor(
			adapter.NewTweetRepository(handler),
		),
		userInteractor: interactor.NewUserInteractor(
			adapter.NewSessionRepository(handler),
			adapter.NewUsersRepository(handler),
		),
		commentInteractor: interactor.NewCommentInteractor(
			adapter.NewCommentRepository(handler),
		),
	}
}

/*
ツイート機能
*/
// CreateTweet ツイートの作成
func (t *PictweetController) CreateTweet(c *gin.Context) {
	ses, err := t.session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	user, err := t.userInteractor.FindBySession(ses)
	if err != nil {
		panic(err)
	}
	text := c.PostForm("text")
	image := c.PostForm("image")
	fmt.Println(text)
	fmt.Println(image)
	if _, err := t.tweetInteractor.Create(user.Id, text, image); err != nil {
		panic(err)
	}

	c.Redirect(http.StatusCreated, "/")
	return

}

// UpdateTweet ツイートの更新
func (t *PictweetController) UpdateTweet(c *gin.Context) {
	ses, err := t.session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	user, err := t.userInteractor.FindBySession(ses)
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

// NewTweet ツイート投稿画面
func (t *PictweetController) NewTweet(c *gin.Context) {
	_, err := t.session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.HTML(http.StatusOK, "newTweet", nil)
}

// EditTweet ツイート編集画面
func (t *PictweetController) EditTweet(c *gin.Context) {
	_, err := t.session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	uuid := c.Query("id")
	tweet, err := t.tweetInteractor.FindByUUID(uuid)
	if err != nil {
		// todo メッセージ：「tweetが見つかりません。」のテンプレート
		return
	}

	fmt.Println(tweet)
	c.HTML(http.StatusOK, "editTweet", gin.H{
		"tweet": tweet,
	})
}

// ReadTweet ツイート詳細画面
func (t *PictweetController) ReadTweet(c *gin.Context) {
	uuid := c.Query("id")
	tweet, err := t.tweetInteractor.FindByUUID(uuid)
	if err != nil {
		// todo tweetが見つかりません的なページに飛ぶ
	}
	comments, err := t.commentInteractor.FindByTweetID(tweet.ID)
	if err != nil {
		// todo ここはnot found のエラーを考える。
	}
	c.HTML(http.StatusOK, "readTweet", gin.H{
		"tweet":    tweet,
		"comments": comments,
	})
}

// DeleteTweet ツイート削除
func (t *PictweetController) DeleteTweet(c *gin.Context) {
	uuid := c.Query("id")
	tweet, err := t.tweetInteractor.FindByUUID(uuid)
	if err != nil {
		// todo tweetはすでに削除されています的なメッセージ
	}
	err = t.tweetInteractor.Delete(tweet)
	if err != nil {
		// todo 削除に失敗しました的なメッセージ
	}
	c.Redirect(http.StatusFound, "/")
}

func (t *PictweetController) Index(c *gin.Context) {
	tweets, err := t.tweetInteractor.Index()
	if err != nil {
		return
	}
	_, err = t.session(c)
	if err != nil {
		c.JSON(http.StatusOK, tweets)
		return
	}
	c.JSON(http.StatusOK, tweets)
	//c.HTML(http.StatusOK, "Index_Private", gin.H{
	//	"tweets": tweets,
	//})

}

/*
コメント機能
*/

// CreateComment コメント作成
func (t *PictweetController) CreateComment(c *gin.Context) {
	ses, err := t.session(c)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}
	user, err := t.userInteractor.FindBySession(ses)
	uuid := c.PostForm("uuid")
	text := c.PostForm("text")
	fmt.Println(uuid)
	fmt.Println(text)
	tweet, err := t.tweetInteractor.FindByUUID(uuid)
	if err != nil {
		// todo このツイートは削除されました的なメッセージ。
	}
	_, err = t.commentInteractor.Create(user.Id, tweet.ID, text)
	if err != nil {
		// todo
	}
	//url := fmt.Sprintf("/tweet/read?id=%s", uuid)
	//c.Redirect(http.StatusFound, url)
}

func (t *PictweetController) session(c *gin.Context) (ses *model.Session, err error) {
	cookie, err := c.Request.Cookie("_cookie")
	fmt.Println(cookie)
	if err == nil {
		ses, err = t.userInteractor.CheckSession("1")
		return
	}
	return &model.Session{
		ID:        1,
		Uuid:      "1",
		Email:     "y@ike",
		UserID:    1,
		CreatedAt: time.Now(),
	}, nil
}
