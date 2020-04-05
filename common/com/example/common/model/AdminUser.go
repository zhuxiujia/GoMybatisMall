package model

import "time"

type AdminUser struct {
	Id         string    `json:"id" gm:"id"`
	Phone      string    `json:"phone"`
	Pwd        string    `json:"pwd"`
	Enable     int       `json:"enable"`
	DeleteFlag int       `json:"delete_flag" gm:"logic"`
	CreateTime time.Time `json:"create_time"`
	Version    int       `json:"version" gm:"version"`
	RealName   string    `json:"real_name"`
	Remark     string    `json:"remark"`
	RoleIds    string    `json:"role_ids"` //json数组
}
