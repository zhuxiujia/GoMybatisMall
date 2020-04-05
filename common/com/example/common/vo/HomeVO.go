package vo

type HomeVO struct {
	BannerPageVO  PageVO `json:"banner_data"`
	NoticePageVO  PageVO `json:"notice_data"`
	ProductPageVO PageVO `json:"product_data"`
}
