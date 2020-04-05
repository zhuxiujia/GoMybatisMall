package utils

import (
	"testing"
	"time"
)

func TestDateId(t *testing.T) {

	println(DateId())
	time.Sleep(time.Nanosecond)
	println(DateId())
	time.Sleep(time.Nanosecond)
	println(DateId())
}
