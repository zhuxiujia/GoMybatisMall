package password

import (
	"github.com/zhuxiujia/crypto/bcrypt"
)

type BCryptPasswordEncoder struct {
}

//加密
func (BCryptPasswordEncoder) Encode(rawPassword string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(rawPassword), 0)
	return string(bytes)
}

//对比密码
func (BCryptPasswordEncoder) Matches(rawPassword string, hashedPassword string) bool {
	var err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	if err != nil {
		return false
	} else {
		return true
	}
}
