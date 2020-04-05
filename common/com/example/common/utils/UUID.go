package utils

import (
	"github.com/google/uuid"
)

func CreateUUID() string {
	// 创建
	uuid := uuid.New()
	var uuidString = uuid.String()
	return uuidString[0:32]
}
