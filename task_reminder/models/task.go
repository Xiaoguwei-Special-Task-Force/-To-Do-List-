package models

import (
	"encoding/json"
	"fmt"
	"os"

	// "reminder/client"
	"reminder/common"
	// "reminder/service"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	// "github.com/sirupsen/logrus"
)

// 新增同步状态枚举
const (
    SyncCreate = iota
    SyncUpdate
    SyncDelete
)

// 任务模型
type Task struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Category    common.TaskCategory   `json:"task_category"` // 新增分类字段
	Deadline    time.Time     `json:"deadline"`
	ReminderMsg string        `json:"reminder_msg"`
	Recipient	string		  `json:"recipient"`
	Completed   bool          `json:"completed"`
	Timer       *time.Timer   `json:"-"`
	mu   		sync.Mutex
	notifyChan chan<- Notification // 通知通道
}

// 修改通知消息结构
type Notification struct {
    Category common.TaskCategory
	Recipient	string
    Message  string
}

// 任务管理器
type TaskManager struct {
	mu       sync.RWMutex
	tasks    map[string]*Task
	storage  string
	notifyCh chan Notification
	wg sync.WaitGroup
}

func NewTaskManager(storagePath string) *TaskManager {
	m := &TaskManager{
		tasks:    make(map[string]*Task),
		storage:  storagePath,
		notifyCh: make(chan Notification, 100),
	}
	m.loadTasks()
	return m
}

// 新增同步方法
func (m *TaskManager) SyncTask(action int, task *Task) {
    // m.mu.Lock()
    // defer m.mu.Unlock()

    switch action {
    case SyncCreate, SyncUpdate:
        // 停止旧定时器
        if old, exists := m.tasks[task.ID]; exists {
            old.mu.Lock()
            if old.Timer != nil {
                old.Timer.Stop()
            }
            old.mu.Unlock()
        }
        
        // 设置新定时器
        task.notifyChan = m.notifyCh
        task.resetReminder()
        m.tasks[task.ID] = task

    case SyncDelete:
        if t, exists := m.tasks[task.ID]; exists {
            t.mu.Lock()
            if t.Timer != nil {
                t.Timer.Stop()
            }
            t.mu.Unlock()
            delete(m.tasks, task.ID)
        }
    }
}

// 任务实例方法
func (t *Task) resetReminder() {
    // t.mu.Lock()
    // defer t.mu.Unlock()

    duration := time.Until(t.Deadline)
    if duration < 0 {
        duration = 0 // 立即触发过期任务
    }

    t.Timer = time.AfterFunc(duration, func() {
        t.notifyChan <- Notification{
            Category:  t.Category,
            Message:   fmt.Sprintf("[%s] 任务到期: %s", t.ID, t.ReminderMsg),
            Recipient: t.Recipient,
        }
    })
	
}

// 添加任务并设置定时器
func (m *TaskManager) AddTask(t *Task)  {
	m.wg.Add(1)
	logrus.Infof("开始同步新增提醒任务>>>>>")
	m.SyncTask(SyncCreate, t)
}

// 新增UpdateTask方法
func (m *TaskManager) UpdateTask(updatedTask *Task) {
	// 
	logrus.Infof("开始同步变更提醒任务>>>>>")
    m.SyncTask(SyncUpdate, updatedTask)
}

// 完成任务
func (m *TaskManager) CompleteTask(id string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if task, exists := m.tasks[id]; exists {
		task.Completed = true
		if task.Timer != nil {
			task.Timer.Stop()
		}
	}
	m.saveTasks()
	m.wg.Done()
}

// 持久化存储
func (m *TaskManager) saveTasks() error {
	data, err := json.MarshalIndent(m.tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(m.storage, data, 0644)
}

func (m *TaskManager) loadTasks() {
	data, err := os.ReadFile(m.storage)
	if err != nil {
		return
	}
	json.Unmarshal(data, &m.tasks)
}

func (m *TaskManager) Wait() {
    m.wg.Wait()
    close(m.notifyCh) // 关闭通知通道
}

// 通知处理器
func (m *TaskManager) StartNotifier(emailSvc *EmailService) {
	go func() {
		for n := range m.notifyCh {
			// 控制台输出保留
            logrus.Infof("\n[%-6s] %s\n", common.ColorCategory(n.Category), n.Message)
            
            // 新增邮件发送
            // go func(notification Notification) {
            //     err := emailSvc.SendNotification(notification)
            //     if err != nil {
            //         logrus.Warn("邮件发送失败: %v | 消息: %s", err, notification.Message)
            //     }
            // }(n)
		}
	}()
}

