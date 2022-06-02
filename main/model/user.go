package model

import "time"

type User struct {
	ID        uint       `gorm:"primaryKey; autoIncrement" json:"id,omitempty"`
	Name      string     `gorm:"not null; size:150" json:"name,omitempty"`
	Email     string     `gorm:"not null; unique; size:300" json:"email,omitempty"`
	Password  string     `gorm:"not null" json:"-"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	Posts     []Post     `gorm:"constraint:OnDelete:CASCADE" json:"posts,omitempty"`
	Follows   []User     `gorm:"many2many:user_follows; constraint:OnDelete:CASCADE" json:"follows,omitempty"`
}

type UserRe struct {
	ID        uint   `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"-"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	Posts     []Post `json:"posts,omitempty"`
	Follows   []User `json:"follows,omitempty"`
}
