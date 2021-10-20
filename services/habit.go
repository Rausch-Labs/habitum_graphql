package services

import (
	"errors"
	"gorm.io/gorm"
	"github.com/suisuss/habitum_graphQL/models"
	"fmt"
	"strings"
)

type HabitServiceI interface {
	CreateHabit(habit *models.Habit) (*models.Habit, error)
	DeleteHabit(id string) (*models.Habit, error)
	GetAllHabitsByName(name string) ([]*models.Habit, error)
	UpdateHabit(habit *models.Habit) (*models.Habit, error)
	GetHabitByID(id string) (*models.Habit, error)
	GetAllHabits() ([]*models.Habit, error)
}


type HabitServiceS struct {
	db *gorm.DB
}

func NewHabitService(db *gorm.DB) *HabitServiceS {
	return &HabitServiceS{
		db,
	}
}



func (s *HabitServiceS) UpdateHabit(habit *models.Habit) (*models.Habit, error) {

	err := s.db.Save(&habit).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return nil, errors.New("message [SOMETHING] already taken")
		}
		return nil, err
	}

	return habit, nil

}

func (s *HabitServiceS) GetHabitByID(id string) (*models.Habit, error) {

	var habit = &models.Habit{}

	err := s.db.Where("id = ?", id).Take(&habit).Error
	if err != nil {
		return nil, err
	}

	return habit, nil
}


func (s *HabitServiceS) CreateHabit(habit *models.Habit) (*models.Habit, error) {

	// Make sure habit name isn't taken
	oldHabits, _ := s.GetAllHabitsByName(habit.Name)
	if len(oldHabits) > 0 {
		for _, oldHabit := range oldHabits {
			if oldHabit.Name == habit.Name {
				return nil, errors.New("That habit name is taken")
			}
		}
	}

	err := s.db.Create(&habit).Error
	if err != nil {
		return nil, err
	}

	return habit, nil
}

func (s *HabitServiceS) DeleteHabit(id string) (*models.Habit, error) {

	habit := &models.Habit{}

	err := s.db.Where("id = ?", id).Delete(&habit).Error

	fmt.Println(err)
	if err != nil {
		return nil, err
	}


	return habit, nil
}

func (s *HabitServiceS) GetAllHabitsByName(name string) ([]*models.Habit, error) {

	var habits []*models.Habit

	err := s.db.Where("name = ?", name).Find(&habits).Error
	if err != nil {
		return nil, err
	}

	return habits, nil

}


func (s *HabitServiceS) GetAllHabits() ([]*models.Habit, error) {

	var habits []*models.Habit

	err := s.db.Table("habits").Find(&habits).Error
	if err != nil {
		return nil, err
	}

	return habits, nil

}
