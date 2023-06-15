package jobs

import (
	"github.com/robfig/cron/v3"
	"github.com/seydinadiouf/anywr-cron/model"
	"gorm.io/gorm"
	"log"
)

func CreateJob(c *cron.Cron, task model.ShedTask, db *gorm.DB) int {
	var entryId cron.EntryID
	var err error
	if task.IsRescheduleJob == true {
		entryId, err = c.AddJob(task.Schedule, DoReScheduleJob(task, db, c))
	} else {
		entryId, err = c.AddJob(task.Schedule, DoJob(task))
	}
	if err != nil {
		log.Fatal(err)
	}
	return int(entryId)

}
