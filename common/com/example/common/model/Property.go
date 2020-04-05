package model

import "time"

type Property struct {
	Id       string `json:"id"`
	UserId   string `json:"user_id"`
	Integral int    `json:"integral"` //积分
	Amount   int    `json:"amount"`   //账户余额

	Version    int       `json:"version"`
	CreateTime time.Time `json:"create_time"`
	DeleteFlag int       `json:"delete_flag"`
}
