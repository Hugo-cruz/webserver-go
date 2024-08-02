package domain

import "time"

type Device struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	Name         string    `gorm:"size:255;not null"`
	Brand        string    `gorm:"size:255;not null"`
	CreationTime time.Time `gorm:"not null"`
}

type DeviceModel struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	Name         string    `gorm:"size:255;not null"`
	Brand        string    `gorm:"size:255;not null"`
	CreationTime time.Time `gorm:"not null"`
}
