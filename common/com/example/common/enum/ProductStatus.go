package enum

import "fmt"

type ProductStatus int

func (it ProductStatus) New(t int) ProductStatus {
	switch t {
	case -1:
		return ProductStatus_Disable
	case 1:
		return ProductStatus_Enable

	}
	panic("不支持产品状态! type=" + fmt.Sprint(t))
}

const (
	ProductStatus_Disable ProductStatus = -1
	ProductStatus_Enable  ProductStatus = 1
)
