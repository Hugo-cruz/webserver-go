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
	common "webserver/src/commom"
)

type SQLiteRepository struct {
	DB *gorm.DB
}

func (d SQLiteRepository) FindByBrand(brand string) ([]domain.Device, error) {
	var devices []domain.Device
	result := d.DB.Where("brand = ?", brand).Find(&devices)
	if result.RowsAffected == 0 {
		return []domain.Device{}, errors.New("device not found")
	}
	return devices, nil
}

func NewSQLiteRepository(db *gorm.DB) *SQLiteRepository {
	return &SQLiteRepository{
		DB: db,
	}
}

func (d SQLiteRepository) Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("urls.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (d SQLiteRepository) Initialize() error {
	sqlFile, err := filepath.Abs("sample.sql")
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
		log.Printf("%t SQL: %v", valid, err)
		return err
	}
	d.DB.Exec(commands)
	fmt.Println("Database initialized successfully")
	return nil
}

func (d SQLiteRepository) Save(device *domain.Device) error {
	err := d.DB.Save(&device)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func (d SQLiteRepository) Update(device *domain.Device) error {
	result := d.DB.Model(&device).Where("id = ?", device.ID).Updates(device)
	if result.RowsAffected == 0 {
		return errors.New("device not updated")
	}
	return nil
}

func (d SQLiteRepository) Delete(deviceId int) error {
	var deviceModel domain.Device
	result := d.DB.Find("id = ?", deviceId).Delete(&deviceModel)
	if result == nil {
		return errors.New("device not found")
	}
	return nil
}

func (d SQLiteRepository) FindById(deviceId int) (*domain.Device, error) {
	var device domain.Device
	result := d.DB.First(&device, deviceId)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(common.ErrDeviceNotFound)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &device, nil
}

func (d SQLiteRepository) FindAll() ([]*domain.Device, error) {
	var Device []*domain.Device
	result := d.DB.Find(&Device)
	if result != nil {
		log.Println(Device)
		return Device, nil
	}
	return Device, errors.New("devices not found")
}

func NewDatabase(db *gorm.DB) Repository {
	return &SQLiteRepository{DB: db}
}

func NewDatabaseConnection(filepath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(filepath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected With Success")
	return db, nil
}

func isSQLValid(sql string) (bool, error) {
	_, err := sqlparser.Parse(sql)
	if err != nil {
		return false, err
	}
	return true, nil
}
