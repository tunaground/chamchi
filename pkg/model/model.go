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
	Threads   []Thread  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Thread struct {
	ID        uint       `json:"id" gorm:"primaryKey";`
	BoardID   uint       `json:"board_id" gorm:"notNull;index:idx_board_title;"`
	Title     string     `json:"title" gorm:"size:50;notNull;index:idx_board_title;"`
	Password  string     `json:"password" gorm:"size:256;notNull;"`
	Status    string     `json:"status" gorm:"size:16;notNull;index;"`
	CreatedAt time.Time  `json:"created_at" gorm:"notNull;autoCreateTime;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"notNull;autoUpdateTime;index;"`
	DeletedAt time.Time  `json:"deleted_at" gorm:"notNull;autoDeleteTime;index;"`
	Responses []Response `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Response struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	ThreadID   uint      `json:"thread_id" gorm:"notNull;index:idx_thread_sequence;"`
	Sequence   uint      `json:"sequence" gorm:"notNull;index:idx_thread_sequence;"`
	Username   string    `json:"username" gorm:"size:60;notNull;index;"`
	UserID     string    `json:"user_id" gorm:"size:10;notNull;index;"`
	IP         string    `json:"ip" gorm:"size:15;notNull;index;"`
	Content    string    `json:"content" gorm:"type:TEXT;size:20000;notNull;"`
	Attachment string    `json:"attachment" gorm:"size:100;notNull;"`
	Youtube    string    `json:"youtube" gorm:"size:100;notNull;"`
	CreatedAt  time.Time `json:"created_at" gorm:"notNull;autoCreateTime;"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"notNull;autoUpdateTime;"`
	DeletedAt  time.Time `json:"deleted_at" gorm:"notNull;autoDeleteTime;"`
}
