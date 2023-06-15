package main

import (
	"github.com/robfig/cron/v3"
	"github.com/seydinadiouf/anywr-cron/config"
	"github.com/seydinadiouf/anywr-cron/model"
	"github.com/seydinadiouf/anywr-cron/task"
	"log"
)

func main() {
	// Connect To Database
	config.DatabaseInit()
	db := config.DB()

	dbGorm, err := db.DB()
	if err != nil {
		panic(err)
	}

	err = dbGorm.Ping()
	if err != nil {
		return
	}

	errorMigration := db.AutoMigrate(&model.ShedTask{})
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", errorMigration)

	}

	c := cron.New()
	c.Start()
	defer c.Stop()

	task.ScheduledTasks(c, db)

	select {}

}
