package repository

import (
	_ "github.com/mattn/go-sqlite3"
	"webserver/src/app/domain"
)

type Repository interface {
	Initialize() error
	Save(device domain.Device) error
	Update(device domain.Device) error
	Delete(deviceId int) error
	FindById(deviceId int) (domain.Device, error)
	FindByBrand(brand string) ([]domain.Device, error)
	FindAll() ([]domain.Device, error)
}
