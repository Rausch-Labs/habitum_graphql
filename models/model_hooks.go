package models

import (
	"gorm.io/gorm"
	"github.com/twinj/uuid"
)

func (mod *Habit) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.NewV4()
	tx.Model(mod).Update("id", uuid.String())
	return
}

func (mod *HabitLog) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.NewV4()
	tx.Model(mod).Update("id", uuid.String())
	return
}

func (mod *User) BeforeCreate(tx *gorm.DB) (err error) {
	uuid := uuid.NewV4()
	tx.Model(mod).Update("id", uuid.String())
	return
}