package model

import "time"

type MallCoverImage struct {
	Id         string    `json:"id" gm:"id"`
	Img        string    `json:"img"`
	SkuId      string    `json:"sku_id"`
	Version    int       `json:"version" gm:"version"`
	CreateTime time.Time `json:"create_time"`
	DeleteFlag int       `json:"delete_flag" gm:"logic"`
}
