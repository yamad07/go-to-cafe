package entity

import "time"

type User struct {
	ID          int64      `gorm:"column:id"`
	AccountID   int64      `gorm:"column:account_id"`
	Name        string     `gorm:"column:name"`
	Description string     `gorm:"column:description"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
}
