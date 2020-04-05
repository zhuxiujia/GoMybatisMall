package model

import "time"

type MallClass struct {
	Id         string    `json:"id" gm:"id"`
	Name       string    `json:"name"`     //分类名称
	LogoImg    string    `json:"logo_img"` //分类logo
	Version    int       `json:"version" gm:"version"`
	CreateTime time.Time `json:"create_time"`
	DeleteFlag int       `json:"delete_flag" gm:"logic"`
}
