package vo

import (
	"math"
)

type Pageable struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}

func (it Pageable) New(page int, size int) Pageable {
	var pageRequest = Pageable{}
	if page < 0 {
		var err = ResultVO{}.NewError(0, "page index must be >= 1 !")
		panic(err)
	} else if size < 1 {
		var err = ResultVO{}.NewError(0, "page size must not be >= 1 !")
		panic(err)
	} else {
		pageRequest.Page = page * size
		pageRequest.PageSize = size
		return pageRequest
	}
}

//继承他必须有属性`json:"content"`
type PageVO struct {
	Pageable
	TotalPages int         `json:"totalPages"`
	Total      int64       `json:"total"`
	Content    interface{} `json:"content"`
}

func (it PageVO) New(pageable Pageable, total int64, content interface{}) PageVO {
	it.SetContent(content)
	it.Pageable = pageable
	it.Total = total
	it.TotalPages = it.countTotalPages(it.PageSize, it.Total)
	return it
}

//func (it *PageVO) DecodeContent(result interface{}) error {
//	//var e = json.Unmarshal(it.Content, result)
//	//if e != nil {
//	//	return e
//	//}
//	return nil
//}

func (it *PageVO) SetContent(result interface{}) {
	//var d, _ = json.Marshal(result)
	//it.Content = d
	it.Content = result
}

func (it *PageVO) countTotalPages(pageSize int, total int64) int {
	if pageSize <= 0 {
		return 1
	} else {
		return int(math.Ceil(float64(total) / float64(pageSize)))
	}
}
