package cronjobs

import (
	"testing"
	"time"
)

/*
type Task struct {
	Id           int
	Name 		 string
	CronSpec     string   //'时间表达式',
	Concurrent   int      //'是否只允许一个实例',
	Command      string  //'命令详情',
	Status       int     //'0停用 1启用',
	Timeout      int     //'超时设置',,
	ExecuteCount int     //'累计执行次数',
	SSH          *ServerSSH  //远程服务器信息
	PrevTime     int64  //'上次执行时间',
	CreateTime   int64  //'创建时间',
	SendMsg      func(log *TaskLog)//发送消息插件

}
*/

/*
func InitJobs() {
	list, _ := models.TaskGetList(1, 1000000, "status", 1)
	for _, task := range list {
		job, err := NewJobFromTask(task)
		if err != nil {
			beego.Error("InitJobs:", err.Error())
			continue
		}
		AddJob(task.CronSpec, job)
	}
}
*/
func Test_cron(t *testing.T) {
	taskList := []*Task{}
	a := &Task{
		Name:       "test",
		CronSpec:   "0/1 * * * * ?",
		Concurrent: 1,
		Command:    "ls -a;sleep 1",
		Timeout:    20,
		Status:     TASK_ENABLE,
	}

	c := &Task{
		Name:       "test12",
		CronSpec:   "0/2 * * * * ?",
		Concurrent: 1,
		Command:    "ls",
		Timeout:    10,
		Status:     TASK_ENABLE,
	}

	taskList = append(taskList, a)
	taskList = append(taskList, c)

	//
	//go func() {
	//	for {
	//		logList := LogDetail(100)
	//		for i := range logList {
	//			fmt.Println("aaaaaaaaaaaaaaaaaaaaaa")
	//			fmt.Println("Error",logList[i].Error)
	//			fmt.Println("Output",logList[i].Output)
	//			fmt.Println("Status",logList[i].Status)
	//			fmt.Println("TaskId",logList[i].TaskId)
	//			fmt.Println("CreateTime",logList[i].CreateTime)
	//			fmt.Println("ProcessTime",logList[i].ProcessTime)
	//			time.Sleep(1 * time.Second)
	//		}
	//	}
	//}()

	for i, b := range taskList {
		if b.Status == TASK_DISABLE {
			continue
		}
		b.Id = i + 1
		job, err := NewJobFromTask(b)
		if err != nil {
			t.Log(err)
		}

		err = AddJob(b.CronSpec, job)
		if err != nil {
			t.Log(err)
		}
	}

	StartJob()

	time.Sleep(5 * time.Second)
	DelJob(2)
	select {}
	//
	//}
	//job.Run()

}
