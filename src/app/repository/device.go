package repository

import (
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
	"webserver/src/app/domain"
)

//go:generate mockgen -source=./device.go -destination=./mock/repository_mock.go -package=repository
type Repository interface {
	Connect() (*gorm.DB, error)
	Initialize() error
	Save(device *domain.Device) error
	Update(device *domain.Device) error
	Delete(deviceId int) error
	FindById(deviceId int) (*domain.Device, error)
	FindByBrand(brand string) ([]domain.Device, error)
	FindAll() ([]*domain.Device, error)
}
