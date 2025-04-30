package main

import (
	// "reminder/models"
	"reminder/service"
	"time"
	// "reminder/test"
)

func main() {
	// test.Remind()
	// test.TimeTick()
	service := service.NewTaskService()
	service.Manager.StartNotifier()

	// 添加测试任务
	service.CreateTask("项目会议", time.Now().Add(7*time.Second), "记得准备会议材料")
	service.CreateTask("代码审查", time.Now().Add(30*time.Second), "检查PR#1234")

	// 保持主线程运行
	select {}
}