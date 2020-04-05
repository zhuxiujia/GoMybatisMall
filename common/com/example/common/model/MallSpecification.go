package model

import "time"

type MallSpecification struct {
	Id         string    `json:"id" gm:"id"`
	SkuId      string    `json:"sku_id"`
	Name       string    `json:"name"` //产品规格名称
	Version    int       `json:"version" gm:"version"`
	CreateTime time.Time `json:"create_time"`
	DeleteFlag int       `json:"delete_flag" gm:"logic"`
}
