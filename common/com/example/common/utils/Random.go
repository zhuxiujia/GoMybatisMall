package utils

func ToDigitString(bytess []byte) string {
	var zero = byte('0')
	var resultBytes = make([]byte, 0)
	for _, item := range bytess {
		resultBytes = append(resultBytes, item+zero)
	}
	return string(resultBytes)
}
