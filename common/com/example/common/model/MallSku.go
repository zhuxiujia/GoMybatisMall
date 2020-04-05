package model

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/enum"
	"time"
)

type MallSku struct {
	Id             string             `json:"id" gm:"id"`
	ClassName      string             `json:"class_name"`       //分类Id
	Title          string             `json:"title"`            //标题
	SecondTitle    string             `json:"second_title"`     //副标题
	Content        string             `json:"content"`          //内容
	TotalNum       int                `json:"total_num"`        //库存产品总数
	RemainNum      int                `json:"remain_num"`       //库存产品剩余数
	ShopAmount     int                `json:"shop_amount"`      //市场价
	Amount         int                `json:"amount"`           //最低价
	OrderTimeLimit int                `json:"order_time_limit"` //下单次数(0则不限制)
	Status         enum.MallSkuStatus `json:"status"`           //产品状态（未上架，已上架，已下架）
	Tag1           string             `json:"tag1"`             //热销
	Tag2           string             `json:"tag2"`             //新品
	Tag3           string             `json:"tag3"`             //精选

	Version    int       `json:"version" gm:"version"`
	CreateTime time.Time `json:"create_time"`
	DeleteFlag int       `json:"delete_flag" gm:"logic"`
}
