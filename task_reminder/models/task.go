package models

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

// 任务模型
type Task struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Deadline    time.Time     `json:"deadline"`
	ReminderMsg string        `json:"reminder_msg"`
	Completed   bool          `json:"completed"`
	Timer       *time.Timer   `json:"-"`
}

// 任务管理器
type TaskManager struct {
	mu       sync.RWMutex
	tasks    map[string]*Task
	storage  string
	notifyCh chan string
}

func NewTaskManager(storagePath string) *TaskManager {
	m := &TaskManager{
		tasks:    make(map[string]*Task),
		storage:  storagePath,
		notifyCh: make(chan string, 100),
	}
	m.loadTasks()
	return m
}

// 添加任务并设置定时器
func (m *TaskManager) AddTask(t *Task) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()
	if t.Deadline.Before(now) {
		return fmt.Errorf("deadline has already passed")
	}

	// 设置定时器
	t.Timer = time.AfterFunc(time.Until(t.Deadline), func() {
		m.notifyCh <- fmt.Sprintf("[ALERT] %s: %s", t.Name, t.ReminderMsg)
		m.CompleteTask(t.ID)
	})

	m.tasks[t.ID] = t
	return m.saveTasks()
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

// 通知处理器
func (m *TaskManager) StartNotifier() {
	go func() {
		for msg := range m.notifyCh {
			fmt.Printf("\n%s\n", msg)
		}
	}()
}