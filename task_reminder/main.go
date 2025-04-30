package main

import (
	// "reminder/models"
	"crypto/tls"
	"reminder/common"
	"reminder/models"
	// "reminder/service"
	"time"
	"github.com/go-mail/mail"
	// "reminder/test"
)

func main() {
	// test.Remind()
	// test.TimeTick()
	// 存在死锁bug-使用waitGroup解决
	// service := service.NewTaskService()
	// service.Manager.StartNotifier()

	// 添加测试任务
	// service.CreateTask("项目会议", time.Now().Add(7*time.Second), "记得准备会议材料")
	// service.CreateTask("代码审查", time.Now().Add(30*time.Second), "检查PR#1234")

	// 添加不同分类的任务
    // service.CreateTask("项目会议", common.Work, 
    //     time.Now().Add(5*time.Second), "准备技术方案")
    // service.CreateTask("体检提醒", common.Personal,
    //     time.Now().Add(10*time.Second), "空腹体检")
    // service.CreateTask("系统维护", common.System,
    //     time.Now().Add(15*time.Second), "执行维护脚本")

	// 保持主线程运行
	// 初始化邮件服务
    emailSvc := &models.EmailService{
        Config: common.EmailConfig{
            Host:     "smtp.example.com",
            Port:     587,
            Username: "notify@company.com",
            Password: "your_password",
            From:     "Task System <notify@company.com>",
        },
        Dialer: mail.NewDialer("smtp.example.com", 587, "notify", "your_password"),
    }
    emailSvc.Dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} // 测试环境跳过证书验证

    // 启动任务管理器
    manager := models.NewTaskManager("tasks.db")
    manager.StartNotifier(emailSvc)
    
    // 添加测试任务
    manager.AddTask(&models.Task{
        ID:       "task-1",
        Category: common.Work,
        ReminderMsg:  "项目评审会议将于15分钟后开始",
        Recipient: "team@company.com", // 新增收件人字段
        Deadline: time.Now().Add(4 * time.Second),
    })
	select {}
}