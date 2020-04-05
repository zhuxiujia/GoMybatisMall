package model

import "time"

//user表
type User struct {
	Id             string `json:"id"`
	Phone          string `json:"phone"`
	Password       string `json:"password"`
	Realname       string `json:"realname"`
	InvitationCode string `json:"invitation_code"`
	Ghost          int    `json:"ghost"`
	TodaySignIn    int    `json:"today_sign_in"`
	ClientType     int    `json:"client_type"`
	Address        string `json:"address"` //地址
	DefAddressId   string `json:"def_address_id"`
	Avatar         string `json:"avatar"` //头像

	Channel     string `json:"channel"`      //注册渠道
	InviterCode string `json:"inviter_code"` //邀请人 邀请码

	Version    int       `json:"version"`
	CreateTime time.Time `json:"create_time"`
	DeleteFlag int       `json:"delete_flag"`
}

func (it *User) ResetImportantInfo() {
	it.Password = "*"
}
