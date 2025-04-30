package models

import (
	"fmt"
	"reminder/common"
	"sync"
	"time"

	"github.com/go-mail/mail"
)

// 邮件服务结构体
type EmailService struct {
    Config  common.EmailConfig
    Dialer  *mail.Dialer
    mu      sync.RWMutex
}

// 邮件发送核心方法
func (es *EmailService) SendNotification(n Notification) error {
    m := mail.NewMessage()
    m.SetHeader("From", es.Config.From)
    m.SetHeader("To", n.Recipient) // 需在Task结构中新增收件人字段
    m.SetHeader("Subject", fmt.Sprintf("[%s] 任务提醒", n.Category))
    
    // 构建HTML内容
    body := fmt.Sprintf(`
        <div style="border-left: 4px solid %s; padding-left: 12px;">
            <h2>%s 类任务提醒</h2>
            <p>%s</p>
            <small>发送时间: %s</small>
        </div>`,
        common.ColorCategory(n.Category), 
        n.Category, 
        n.Message,
        time.Now().Format("2006-01-02 15:04"),
    )
    m.SetBody("text/html", body)

    // 异步发送
    es.mu.Lock()
    defer es.mu.Unlock()
    return es.Dialer.DialAndSend(m)
}