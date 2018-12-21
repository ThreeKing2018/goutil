package cronjobs

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"time"

	"github.com/ThreeKing2018/goutil/golog"
	"github.com/astaxie/beego"
)

const (
	TASK_DISABLE = 0
	TASK_ENABLE  = 1
)

type Task struct {
	Id           int
	Name         string
	CronSpec     string             //'时间表达式',
	Concurrent   int                //'是否只允许一个实例',
	Command      string             //'命令详情',
	Status       int                //'0停用 1启用',
	Timeout      int                //'超时设置',,
	ExecuteCount int                //'累计执行次数',
	SSH          *ServerSSH         //远程服务器信息
	PrevTime     int64              //'上次执行时间',
	CreateTime   int64              //'创建时间',
	SendMsg      func(log *TaskLog) //发送消息插件

}

type ServerSSH struct {
	Addr   string
	Port   int
	User   string
	Passwd string
	Remark string
}

func NewJobFromTask(task *Task) (*Job, error) {
	if task.Id < 1 {
		return nil, errors.New("ToJob: 缺少id")
	}

	//本地程序执行
	if task.SSH == nil {
		job := NewCommandJob(task.Id, task.Name, task.Command)
		job.task = task
		job.Concurrent = task.Concurrent == 1
		return job, nil
	}

	//远程服务器
	return nil, nil
}

func NewCommandJob(id int, name string, command string) *Job {
	job := &Job{
		id:   id,
		name: name,
	}
	job.runFunc = func(timeout time.Duration) (string, string, error, bool) {
		bufOut := new(bytes.Buffer)
		bufErr := new(bytes.Buffer)
		cmd := exec.Command("sh", "-c", command)
		cmd.Stdout = bufOut
		cmd.Stderr = bufErr
		cmd.Start()
		err, isTimeout := runCmdWithTimeout(cmd, timeout)

		return bufOut.String(), bufErr.String(), err, isTimeout
	}
	return job
}

func runCmdWithTimeout(cmd *exec.Cmd, timeout time.Duration) (error, bool) {
	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	var err error
	select {
	case <-time.After(timeout):
		golog.Warnf("任务执行时间超过%d秒，进程将被强制杀掉: %d", int(timeout/time.Second), cmd.Process.Pid)
		go func() {
			<-done // 读出上面的goroutine数据，避免阻塞导致无法退出
		}()
		if err = cmd.Process.Kill(); err != nil {
			golog.Errorf("进程无法杀掉: %d, 错误信息: %s", cmd.Process.Pid, err)
			beego.Error(fmt.Sprintf("进程无法杀掉: %d, 错误信息: %s", cmd.Process.Pid, err))
		}
		return err, true
	case err = <-done:
		return err, false
	}
}
