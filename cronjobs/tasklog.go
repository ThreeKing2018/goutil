package cronjobs

import (
	"sync"

	"github.com/pkg/errors"
)

var defaultMaxLogLen = 100

const (
	TASKLOG_SUCCESS = 0  // 任务执行成功
	TASKLOG_ERROR   = -1 // 任务执行出错
	TASKLOG_TIMEOUT = -2 // 任务执行超时
)

type TaskLog struct {
	TaskId      int
	Output      string
	Error       string
	Status      int
	ProcessTime int
	CreateTime  int64
}

type TaskLogMana struct {
	mu      sync.Mutex
	TaskLog []*TaskLog
	Count   int
}

func (t *TaskLogMana) Add(log *TaskLog) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.Count >= defaultMaxLogLen {
		return errors.New("超过最大存储的日志数量")
	}
	t.TaskLog = append(t.TaskLog, log)
	t.Count++

	return nil
}

func (t *TaskLogMana) Detail(num int) []*TaskLog {
	if num > defaultMaxLogLen {
		num = defaultMaxLogLen
	}

	if num > len(t.TaskLog) {
		num = len(t.TaskLog)
	}
	return t.TaskLog[:num]
}
