package jobs

import (
	"fmt"
	"github.com/seydinadiouf/anywr-cron/model"
)

type Job struct {
	Task model.ShedTask
}

func DoJob(task model.ShedTask) *Job {
	return &Job{
		Task: task,
	}
}

func (job *Job) Run() {
	fmt.Printf("Executing task: %s\n", job.Task.TaskName)
}
