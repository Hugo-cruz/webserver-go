package repository

import (
	"errors"
	"fmt"
	"github.com/xwb1989/sqlparser"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"webserver/src/app/domain"
)

type DatabaseRepository struct {
	DB *gorm.DB
}

func (d DatabaseRepository) FindByBrand(brand string) ([]domain.Device, error) {
	var devices []domain.Device
	result := d.DB.Where("brand = ?", brand).Find(&devices)
	if result.RowsAffected == 0 {
		return []domain.Device{}, errors.New("device not found")
	}
	return devices, nil
}

func NewDatabaseRepository(db *gorm.DB) *DatabaseRepository {
	return &DatabaseRepository{
		DB: db,
	}
}

func (d DatabaseRepository) Initialize() error {
	sqlFile, err := filepath.Abs("sample.sql")
	fmt.Println(sqlFile)
	if err != nil {
		log.Println(err)
		return err
	}
	fileContent, err := os.ReadFile(sqlFile)
	if err != nil {
		log.Println(err)
		return err
	}
	commands := string(fileContent)
	valid, err := isSQLValid(commands)
	if err != nil {
		log.Println("%s SQL: %v", valid, err)
		return err
	}
	d.DB.Exec(commands)
	fmt.Println("Database initialized successfully")
	return nil
}

func (d DatabaseRepository) Save(device domain.Device) error {
	err := d.DB.Save(&device)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (d DatabaseRepository) Update(device domain.Device) error {
	result := d.DB.Model(&device).Where("id = ?", device.ID).Updates(device)
	if result.RowsAffected == 0 {
		return errors.New("device not updated")
	}
	return nil
}

func (d DatabaseRepository) Delete(deviceId int) error {
	var deviceModel domain.DeviceModel
	result := d.DB.Find("id = ?", deviceId).Delete(&deviceModel)
	if result == nil {
		return errors.New("device not found")
	}
	return nil
}

func (d DatabaseRepository) FindById(deviceId int) (domain.Device, error) {
	var device domain.Device
	result := d.DB.Find(&device, deviceId)
	if result != nil {
		return device, nil
	}
	return domain.Device{}, errors.New("device not found")
}

func (d DatabaseRepository) FindAll() ([]domain.Device, error) {
	var Device []domain.Device
	result := d.DB.Find(&Device)
	if result != nil {
		log.Println(Device)
		return Device, nil
	}
	return Device, errors.New("devices not found")
}

func NewDatabase(db *gorm.DB) Repository {
	return &DatabaseRepository{DB: db}
}

func NewDatabaseConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./data.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("connected")
	return db, nil
}

func isSQLValid(sql string) (bool, error) {
	_, err := sqlparser.Parse(sql)
	if err != nil {
		return false, err
	}
	return true, nil
}
