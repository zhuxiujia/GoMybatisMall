package utils

import (
	"time"
)

func CreateRechargeTime(createTime time.Time, TotalPayTime int, month int) string {
	if TotalPayTime == 1 {
		return createTime.AddDate(0, 1, 0).Format("2006-01-02")
	} else {
		return createTime.AddDate(0, month, 0).Format("2006-01-02")
	}
}
