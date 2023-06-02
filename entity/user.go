package entity

import "time"

// User 사용자 entity
type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// TableName pluralize
func (User) TableName() string {
	return "users"
}
