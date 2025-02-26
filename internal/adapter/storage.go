package adapter

import (
	"ScheduleAssist/internal/logger"
	"ScheduleAssist/internal/model/orm"
	"errors"
	"gorm.io/gorm"
)

func SetUserWithTID(id string) (*orm.User, error) {
	user, err := GetUserByTID(id)
	if err != nil {
		logger.Error("failed to get user: %v", err)
		return nil, err
	}
	if user != nil {
		return user, nil
	}
	user = &orm.User{
		TelegramID: &id,
	}
	result := GetDB().Create(user)
	if result.Error != nil {
		logger.Error("failed to save user: %v", result.Error)
		return nil, result.Error
	}
	logger.Debug("Row saved %d", result.RowsAffected)
	return user, nil
}

func GetUserByTID(id string) (*orm.User, error) {
	var user orm.User
	tx := GetDB().First(&user, "telegram_id=?", id)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if tx.Error != nil {
		logger.Error("failed to get user: %v", tx.Error)
		return nil, tx.Error
	}
	return &user, nil
}

func CreatTasks(value *[]orm.Task) error {
	result := GetDB().Create(value)
	if result.Error != nil {
		logger.Error("failed to save value: %v", result.Error)
		return result.Error
	}
	logger.Debug("Row saved %d", result.RowsAffected)
	return nil
}
