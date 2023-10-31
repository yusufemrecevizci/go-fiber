package models

import "time"

type Product struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
}
