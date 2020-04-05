package model

import "time"

type KV struct {
	Id         string    `json:"id"`
	Value      string    `json:"value"`
	Remark     string    `json:"remark"`
	CreateTime time.Time `json:"create_time"`
}
