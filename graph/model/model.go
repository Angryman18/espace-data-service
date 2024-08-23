package model

type UserData struct {
	UserId   string `bson:"id"`
	Email    string `bson:"email"`
	Fullname string `bson:"fullname"`
	Avatar   string `bson:"avatar"`
	Id       string `bson:"_id"`
	Provider string `bson:"provider"`
	Profile  string `bson:"profile"`
	Username string `bson:"username"`
}
