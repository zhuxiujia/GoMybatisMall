package core_util

import "testing"

func TestAmountUtil_wx(test *testing.T) {
	println(ToThirdPayAmount("wx", 13500.0))
}

func TestAmountUtil_zfb(test *testing.T) {
	println(ToThirdPayAmount("zfb", 1350000.0))
}
