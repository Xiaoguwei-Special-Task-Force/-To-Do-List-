package test

import (
	"fmt"
	"time"
)

func TimeTick() {
	// 设置倒计时时长（秒）
	countdown := 10

	// 创建定时器（总时长）和节拍器（每秒更新）
	timer := time.NewTimer(time.Duration(countdown) * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	defer timer.Stop()

	fmt.Println("倒计时开始：")
	
	// 实时显示剩余时间
	go func() {
		for remaining := countdown; remaining > 0; remaining-- {
			// \r 实现行内刷新，%2d 保持两位数格式
			fmt.Printf("\r剩余时间: %2d 秒", remaining)
			<-ticker.C // 等待下一次 tick
		}
	}()

	// 等待倒计时结束
	<-timer.C
	fmt.Printf("\r任务提醒：时间到！%s\n", time.Now().Format("15:04:05"))
}