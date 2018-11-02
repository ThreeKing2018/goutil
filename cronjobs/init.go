package cronjobs

import (
	"github.com/ThreeKing2018/goutil/cronjobs/cron"
)

var _cron Croner
var _TaskLogMana *TaskLogMana

func init() {
	_cron = NewCronJob()
	_TaskLogMana = &TaskLogMana{
		Count: 0,
	}
}

func AddJob(spec string, job *Job) error {
	return _cron.AddJob(spec, job)
}

func DelJob(jobID int) {
	_cron.DelJob(jobID)
}

func DetailJob() []*cron.Entry {
	return _cron.DetailJob()
}

func StartJob() {
	_cron.StartJob()
}

func StopJob() {
	_cron.StopJob()
}

func LogDetail(num int) []*TaskLog {
	return _TaskLogMana.Detail(num)
}
