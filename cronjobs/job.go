package cronjobs

import (
	"time"

	"github.com/ThreeKing2018/goutil/golog"
)

type Job struct {
	id         int                                               // 任务ID
	name       string                                            // 任务名称
	task       *Task                                             // 任务对象
	runFunc    func(time.Duration) (string, string, error, bool) // 执行函数
	status     int                                               // 任务状态，大于0表示正在执行中 	// 上一次运行状态
	Concurrent bool                                              // 同一个任务是否允许并行执行
}

func (j *Job) Run() {
	if j.Concurrent && j.status > 0 {
		golog.Warnw("任务上一次执行尚未结束，本次被忽略。",
			"jobid", j.id)
	}

	defer func() {
		//if err := recover(); err != nil {
		//	golog.Errorw("--recover-- ", "err",err)
		//}
	}()

	golog.Debugw("开始执行任务", "jobid", j.id)

	j.status++
	defer func() {
		j.status--
	}()

	//默认超时时间 1小时
	t := time.Now()
	var timeout time.Duration

	if j.task.Timeout > 0 {
		timeout = time.Second * time.Duration(j.task.Timeout)
	} else {
		timeout = time.Duration(time.Hour)
	}

	cmdOut, cmdErr, err, isTimeout := j.runFunc(timeout)
	ut := time.Now().Sub(t) / time.Millisecond

	//记录日志
	log := new(TaskLog)
	log.Output = cmdOut
	log.Error = cmdErr
	log.ProcessTime = int(ut)
	log.CreateTime = t.Unix()
	log.Status = TASKLOG_SUCCESS

	if isTimeout { //任务执行超时了
		log.Status = TASKLOG_TIMEOUT

	} else if err != nil { //任务执行出错了
		log.Status = TASKLOG_ERROR

	}

	// 更新日志
	_TaskLogMana.Add(log)
	//更新上次执行时间
	j.task.PrevTime = t.Unix()
	j.task.ExecuteCount++

	//发送消息通知
	if log.Status != TASKLOG_SUCCESS && j.task.SendMsg != nil {
		j.task.SendMsg(log)
	}

}
