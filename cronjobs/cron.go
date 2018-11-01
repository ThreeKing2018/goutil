package cronjobs

import (
	"github.com/robfig/cron"
	"sync"
)

type Croner interface {
	JobAdd(spec string, job *Job) error
	JobStop(jobID int) error
	//JobDel(jobId int) error
	//JobDetail(uuid ...string) error
}

type cronjob struct {
	cron *cron.Cron
	mu   sync.Mutex
}

func NewCronJob() Croner {
	c := &cronjob{
		cron: cron.New(),
	}
	c.cron.Start()
	return c
}

func (c *cronjob) JobAdd(spec string, job *Job) error {
	c.mu.Lock()
	defer c.mu.Lock()

	if c.getEntryById(job.id) != nil {
		return ERR_AlreadyExisted
	}
	return c.cron.AddJob(spec, job)
}

func (c *cronjob) JobStop(jobID int) error {
	c.cron.AddFunc()
	c.cron.
		c.cron.Stop()
	return nil
}

func (c *cronjob) JobDel(jobId int) error {
	return nil
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
