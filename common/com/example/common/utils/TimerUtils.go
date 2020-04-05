package utils

import (
	"time"
)

type StartType int

const (
	StartType_Wait StartType = iota
	StartType_Now
)

type ExecuteType int

const (
	Execute_Default ExecuteType = iota
	Execute_coroutine
)

//启动Timer,调用它请不要使用 go 关键字,内部已经使用协程处理
func StartTimer(startType StartType, exeType ExecuteType, d time.Duration, fn func()) {
	if startType == StartType_Now {
		fn()
	}
	if exeType == Execute_Default {
		run(fn, d)
	} else {
		//启动一条goroutine 处理轮询任务
		go run(fn, d)
	}
}
func run(fn func(), d time.Duration) {
	ticker := time.NewTicker(d)
	for _ = range ticker.C {
		fn()
	}
}
