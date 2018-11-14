package cronjobs

import (
	"github.com/ThreeKing2018/goutil/cronjobs/cron"
	"sync"
)

type Croner interface {
	AddJob(spec string, job *Job) error
	DelJob(jobID int)
	DetailJob() []*cron.Entry
	StartJob()
	StopJob()
}

type cronjob struct {
	cron *cron.Cron
	mu   sync.Mutex
}

func NewCronJob() Croner {
	c := &cronjob{
		cron: cron.New(),
	}
	//c.cron.Start()
	return c
}

func (c *cronjob) AddJob(spec string, job *Job) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.getEntryById(job.id) != nil {
		return ERR_AlreadyExisted
	}
	return c.cron.AddJob(spec, job)
}

func (c *cronjob) DelJob(jobId int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fn := func(e *cron.Entry) bool {
		if v, ok := e.Job.(*Job); ok {
			if v.id == jobId {
				return true
			}
		}
		return false
	}

	c.cron.RemoveJob(fn)
}

func (c *cronjob) DetailJob() []*cron.Entry {
	ret := c.cron.Entries()
	return ret

}

func (c *cronjob) StartJob() {
	c.cron.Start()
}

func (c *cronjob) StopJob() {
	c.cron.Stop()
}

func (c *cronjob) getEntryById(id int) *cron.Entry {
	entries := c.cron.Entries()
	for _, e := range entries {
		if v, ok := e.Job.(*Job); ok {
			if v.id == id {
				return e
			}
		}
	}
	return nil
}
