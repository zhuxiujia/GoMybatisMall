package utils

import (
	"strconv"
	"time"
)

func DateId() string {
	var d = time.Now().UnixNano()
	return strconv.FormatInt(d, 10)
}

func DateIdInt64() int64 {
	var id = DateId()
	result, _ := strconv.ParseInt(id, 10, 64)
	return result
}
