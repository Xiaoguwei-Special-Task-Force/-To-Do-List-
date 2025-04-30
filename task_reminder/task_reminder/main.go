package main

import (
	// "reminder/models"
	"reminder/common"
	"reminder/service"
	"time"
	// "reminder/test"
)

func main() {
	// test.Remind()
	// test.TimeTick()
	// 存在死锁bug-使用waitGroup解决
	service := service.NewTaskService()
	service.Manager.StartNotifier()

	// 添加测试任务
	// service.CreateTask("项目会议", time.Now().Add(7*time.Second), "记得准备会议材料")
	// service.CreateTask("代码审查", time.Now().Add(30*time.Second), "检查PR#1234")

	// 添加不同分类的任务
    service.CreateTask("项目会议", common.Work, 
        time.Now().Add(5*time.Second), "准备技术方案")
    service.CreateTask("体检提醒", common.Personal,
        time.Now().Add(10*time.Second), "空腹体检")
    service.CreateTask("系统维护", common.System,
        time.Now().Add(15*time.Second), "执行维护脚本")

	// 保持主线程运行
	select {}
}