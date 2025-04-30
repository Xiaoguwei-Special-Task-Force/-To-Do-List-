package test

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func Remind() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Println("请输入提醒间隔（秒）：")
	input, _ := reader.ReadString('\n')
	input = input[:len(input)-1] // 去除换行符
	
	seconds, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("输入无效，请输入数字")
		return
	}
	
	duration := time.Duration(seconds) * time.Second
	fmt.Printf("提醒将每 %v 触发一次\n", duration)
	fmt.Println("输入'stop'停止程序")
	
	ticker := time.NewTicker(duration)
	defer ticker.Stop()
	
	go func() {
		for {
			<-ticker.C
			fmt.Println("提醒：时间到了！", time.Now().Format("15:04:05"))
		}
	}()
	
	// 等待用户输入stop
	for {
		cmd, _ := reader.ReadString('\n')
		if cmd == "stop\n" {
			fmt.Println("程序已停止")
			return
		}
	}
}