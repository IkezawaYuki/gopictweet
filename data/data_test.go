package data

var users = []User{
	{
		Name: "Ikezawa Yuki",
		Email: "ikezawa@hogehoge.com"
		Password: "aiueo12345"
	},
	{
		Name:"Taro Yamada",
		Email:"hogehoge@hogehogehoge.com",
		Password:"12345678"
	},
}

func setup(){
	TweetDeleteAll()
	SessionDeleteAll()
	UserDeleteAll()
}