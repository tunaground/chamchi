package model

import (
	"time"
)

type Board struct {
	ID        uint      `json:"id" gorm:"primaryKey;"`
	Key       string    `json:"key" gorm:"size:16;notNull;uniqueIndex;"`
	Name      string    `json:"name" gorm:"size:16;notNull;"`
	CreatedAt time.Time `json:"created_at" gorm:"notNull;autoCreateTime;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"notNull;autoUpdateTime;"`
	DeletedAt time.Time `json:"deleted_at" gorm:"notNull;autoDeleteTime;"`
	Threads   []Thread  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Thread struct {
	ID        uint `gorm:"primaryKey"`
	BoardID   uint
	Title     string    `gorm:"size:50;notNull;index;"`
	Password  string    `gorm:"size:256;notNull;"`
	UserName  string    `gorm:"size:60;notNull;"`
	CreatedAt time.Time `gorm:"notNull;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"notNull;autoUpdateTime;"`
	DeletedAt time.Time `gorm:"notNull;autoDeleteTime;"`
}
