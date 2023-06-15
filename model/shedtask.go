package model

type ShedTask struct {
	ID              uint   `gorm:"primary_key"`
	Schedule        string `gorm:"column:schedule"`
	TaskName        string `gorm:"column:task_name"`
	TaskParam       string `gorm:"column:task_param"`
	EntryId         int    `gorm:"column:entry_id"`
	IsRescheduleJob bool   `gorm:"column:is_reschedule_job"`
}
