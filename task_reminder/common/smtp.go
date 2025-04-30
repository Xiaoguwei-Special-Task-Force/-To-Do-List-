package common

import (
	// "fmt"
	// "reminder/models"
	// "sync"
	// "time"

	// "github.com/go-mail/mail"
)

// 新增邮件配置结构体
type EmailConfig struct {
    Host         string
    Port         int
    Username     string
    Password     string
    From         string
    TemplatePath string // HTML模板路径（可选）
}



