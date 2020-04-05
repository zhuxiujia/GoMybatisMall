package model

import "time"

type PayCallBack struct {
	Id         string    `json:"id" gm:"id"`
	Data       string    `json:"data"`
	CreateTime time.Time `json:"create_time"`
	DeleteFlag int       `json:"delete_flag" gm:"logic"`
}
