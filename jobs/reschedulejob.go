package jobs

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/seydinadiouf/anywr-cron/model"
	"gorm.io/gorm"
)

type ReScheduleJob struct {
	Task model.ShedTask
	DB   *gorm.DB
	Cron *cron.Cron
}

func DoReScheduleJob(task model.ShedTask, db *gorm.DB, cron *cron.Cron) *ReScheduleJob {
	return &ReScheduleJob{
		Task: task,
		DB:   db,
		Cron: cron,
	}
}

func (job *ReScheduleJob) Run() {
	// Stop the existing cron jobs
	job.Cron.Stop()

	// Remove removed task
	entries := job.Cron.Entries()
	for _, entry := range entries {
		var tasks []model.ShedTask
		job.DB.Where("entry_id = ? ", int(entry.ID)).Find(&tasks)
		if len(tasks) == 0 {
			job.Cron.Remove(entry.ID)
		}
	}

	// Add new added task
	var newTaskList []model.ShedTask
	job.DB.Order("id").Where("entry_id IS NULL").Find(&newTaskList)

	for _, currentTask := range newTaskList {
		currentTask.EntryId = CreateJob(job.Cron, currentTask, job.DB)
		if err := job.DB.Save(&currentTask).Error; err != nil {
			fmt.Printf("Failed to update task with ID %d: %s\n", currentTask.ID, err)
			return
		}
	}

	// Restart the cron scheduler
	job.Cron.Start()

	fmt.Println("Task rescheduled")

}
