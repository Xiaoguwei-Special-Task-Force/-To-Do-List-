package service

import (
	"fmt"
	"reminder/models"
	"time"

	"github.com/sirupsen/logrus"
)

type TaskService struct {
	Manager *models.TaskManager
}

func NewTaskService() *TaskService {
	return &TaskService{
		Manager: models.NewTaskManager("tasks.db"),
	}
}

func (s *TaskService) CreateTask(name string, deadline time.Time, msg string) {
	task := &models.Task{
		ID:          fmt.Sprintf("%d", time.Now().UnixNano()),
		Name:        name,
		Deadline:    deadline,
		ReminderMsg: msg,
	}
	logrus.Infof("%s 将在%v之后截止",name,deadline)
	s.Manager.AddTask(task)
}