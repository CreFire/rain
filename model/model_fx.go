package model

type user struct {
	ID int32
	Avatar string
	Description string
	Email string
	MfaKey string
	Nickname string
	Password string
	Username string
	ExpireTime int64
	CreateTime int64
	UpdateTime int64
}

func NewUser(ID int32, avatar string, description string, email string, mfaKey string, nickname string, password string, username string, expireTime int64) *user {
	return &user{ID: ID, Avatar: avatar, Description: description, Email: email, MfaKey: mfaKey, Nickname: nickname, Password: password, Username: username, ExpireTime: expireTime}
}
