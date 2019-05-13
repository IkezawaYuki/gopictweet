package data

var users = []User{
	{
		Nickname: "Ikezawa Yuki",
		Email:    "ikezawa@hogehoge.com",
		Password: "aiueo12345",
	},
	{
		Nickname: "Taro Yamada",
		Email:    "hogehoge@hogehogehoge.com",
		Password: "12345678",
	},
}

func setup() {
	// TweetDeleteAll()
	SessionDeleteAll()
	UserDeleteAll()
}
