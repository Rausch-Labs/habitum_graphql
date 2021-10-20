package services

import (
	"errors"
	"gorm.io/gorm"
	"github.com/suisuss/habitum_graphQL/models"
	"fmt"
	"strings"
)

type HabitLogServiceI interface {
	CreateHabitLog(habitLog *models.HabitLog) (*models.HabitLog, error)
	DeleteHabitLog(id string) (*models.HabitLog, error)
	GetAllHabitLogsByDate(name string) ([]*models.HabitLog, error)
	GetHabitLogByID(id string) (*models.HabitLog, error)
	UpdateHabitLog(habitLog *models.HabitLog) (*models.HabitLog, error)
	GetAllHabitLogs() ([]*models.HabitLog, error)
}


type HabitLogServiceS struct {
	db *gorm.DB
}

func (s *HabitLogServiceS) UpdateHabitLog(habitLog *models.HabitLog) (*models.HabitLog, error) {
	err := s.db.Save(&habitLog).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("message [SOMETHING] already taken")
		}
		return nil, err
	}

	return habitLog, nil
}

func NewHabitLogService(db *gorm.DB) *HabitLogServiceS {
	return &HabitLogServiceS{
		db,
	}
}

func (s *HabitLogServiceS) CreateHabitLog(habitLog *models.HabitLog) (*models.HabitLog, error) {

	// Make sure habit name isn't taken
	oldHabitLogs, _ := s.GetAllHabitLogsByDate(habitLog.Date)
	if len(oldHabitLogs) > 0 {
		for _, oldHabit := range oldHabitLogs {
			if oldHabit.Type == habitLog.Type {
				return nil, errors.New("That habit type has already been logged for that date")
			}
		}
	}

	err := s.db.Create(&habitLog).Error
	if err != nil {
		return nil, err
	}

	return habitLog, nil
}

func (s *HabitLogServiceS) DeleteHabitLog(id string) (*models.HabitLog, error) {

	habitLog := &models.HabitLog{}

	err := s.db.Where("id = ?", id).Delete(&habitLog).Error

	fmt.Println(err)
	if err != nil {
		return nil, err
	}


	return habitLog, nil
}

func (s *HabitLogServiceS) GetAllHabitLogsByDate(date string) ([]*models.HabitLog, error) {

	var habitLogs []*models.HabitLog

	err := s.db.Where("date = ?", date).Find(&habitLogs).Error
	if err != nil {
		return nil, err
	}

	return habitLogs, nil

}

func (s *HabitLogServiceS) GetHabitLogByID(id string) (*models.HabitLog, error) {

	var habitLog = &models.HabitLog{}

	err := s.db.Where("id = ?", id).Take(&habitLog).Error
	if err != nil {
		return nil, err
	}

	return habitLog, nil
}



func (s *HabitLogServiceS) GetAllHabitLogs() ([]*models.HabitLog, error) {

	var habitLogs []*models.HabitLog

	err := s.db.Table("habit_logs").Find(&habitLogs).Error
	if err != nil {
		return nil, err
	}

	return habitLogs, nil

}