package model

import "time"

type PropertyRecord struct {
	Id         string    `json:"id" gm:"id"`
	Amount     int       `json:"amount"`
	CreateTime time.Time `json:"create_time"`
	DeleteFlag int       `json:"delete_flag" gm:"logic"`
}
