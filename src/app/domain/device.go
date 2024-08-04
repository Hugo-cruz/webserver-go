package domain

import "time"

type CreateDevice struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	Name         string    `gorm:"size:255;not null"`
	Brand        string    `gorm:"size:255;not null"`
	CreationTime time.Time `gorm:"not null"`
}

type Device struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	Name         string    `gorm:"size:255;not null"`
	Brand        string    `gorm:"size:255;not null"`
	CreationTime time.Time `gorm:"not null"`
}

func NewMockDevice() *Device {
	return &Device{
		ID:    1,
		Name:  "Teste",
		Brand: "Brand",
	}
}
