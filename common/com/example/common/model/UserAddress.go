package model

import "time"

type UserAddress struct {
	Id            string `json:"id" gm:"id"`
	UserId        string `json:"user_id"`
	RealName      string `json:"real_name"`
	Phone         string `json:"phone"`
	AddressDetail string `json:"address_detail"`

	Version    int       `json:"version" gm:"version"`
	CreateTime time.Time `json:"create_time"`
	DeleteFlag int       `json:"delete_flag" gm:"logic"`
}
