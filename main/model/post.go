package model

import "time"

type Post struct {
	ID        uint       `gorm:"primaryKey; autoIncrement" json:"id,omitempty"`
	Comment   string     `gorm:"not null; size:600" json:"comment,omitempty"`
	UserID    uint       `gorm:"not null;" json:"userid,omitempty"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"createdAt"`
	Writer    User       `gorm:"foreignKey:UserID" json:"writer,omitempty"`
}

type PostRe struct {
	ID        uint   `json:"id,omitempty"`
	Comment   string `json:"comment,omitempty"`
	UserID    uint   `json:"userid,omitempty"`
	UpdatedAt string `json:"updatedAt"`
	CreatedAt string `json:"createdAt"`
	Writer    UserRe `json:"writer,omitempty"`
}
