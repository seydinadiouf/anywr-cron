package task

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/seydinadiouf/anywr-cron/jobs"
	"github.com/seydinadiouf/anywr-cron/model"
	"gorm.io/gorm"
)

func ScheduledTasks(c *cron.Cron, db *gorm.DB) {
	var shedTasks []model.ShedTask
	db.Order("id").Find(&shedTasks)
	for _, task := range shedTasks {
		task.EntryId = jobs.CreateJob(c, task, db)
		if err := db.Save(&task).Error; err != nil {
			fmt.Printf("Failed to update task with ID %d: %s\n", task.ID, err)
			return
		}
	}

}
