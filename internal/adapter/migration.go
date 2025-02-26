package adapter

import "ScheduleAssist/internal/model/orm"

func migration() {
	GetDB().AutoMigrate(&orm.User{}, &orm.Task{})
}
