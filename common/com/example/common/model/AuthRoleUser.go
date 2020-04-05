package model

import "time"

type AuthRole struct {
	Id          string    `json:"id" gm:"id"`
	Name        string    `json:"name"`
	ResourceIds string    `json:"resource_ids"`
	Version     int       `json:"version" gm:"version"`
	CreateTime  time.Time `json:"create_time"`
	DeleteFlag  int       `json:"delete_flag" gm:"logic"`
}
