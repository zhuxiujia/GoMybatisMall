package vo

import (
	"github.com/zhuxiujia/GoMybatisMall/common/com/example/common/model"
	"math"
)

type PayCallBackPageVO struct {
	Pageable
	TotalPages int                 `json:"totalPages"`
	Total      int64               `json:"total"`
	Content    []model.PayCallBack `json:"content"`
}

func (it PayCallBackPageVO) New(pageable Pageable, total int64, content []model.PayCallBack) PayCallBackPageVO {
	it.Content = content
	it.Pageable = pageable
	it.Total = total
	it.TotalPages = it.countTotalPages(it.PageSize, it.Total)
	return it
}
func (it *PayCallBackPageVO) countTotalPages(pageSize int, total int64) int {
	if pageSize <= 0 {
		return 1
	} else {
		return int(math.Ceil(float64(total) / float64(pageSize)))
	}
}
