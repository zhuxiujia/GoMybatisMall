package utils

import (
	"crypto/md5"
	"strconv"
)

const hash_len = 4

func CountHash4(v string) string {
	hash := md5.New()
	hash.Write([]byte(v))
	result := hash.Sum(nil)[0:hash_len]
	var str string
	for i := 0; i < hash_len; i++ {
		str += strconv.Itoa(int(result[i]))
	}
	return str
}
