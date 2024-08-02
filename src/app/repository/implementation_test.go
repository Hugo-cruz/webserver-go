package repository

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
	"testing"
	"webserver/src/app/domain"
)

type RepositoryTestSuite struct {
	suite.Suite
	context.Context
	repository         Repository
	DBConnectionString string
}

func NewMockDatabase() *gorm.DB {
	filepath := ":memory:"
	db, err := NewDatabaseConnection(filepath)
	if err != nil {
		return nil
	}
	if db == nil {
		return nil
	}
	err = db.Exec(`CREATE TABLE IF NOT EXISTS Devices (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name VARCHAR(255) NOT NULL,
        brand VARCHAR(255) NOT NULL,
        creation_time DATETIME NOT NULL
    )`).Error
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
		return nil
	}

	// Insert sample data into the Devices table
	err = db.Exec(`INSERT INTO Devices (name, brand, creation_time)
    VALUES
        ('Device1', 'BrandA', '2024-01-01 10:00:00'),
        ('Device2', 'BrandB', '2024-02-01 11:00:00'),
        ('Device3', 'BrandC', '2024-03-01 12:00:00')`).Error
	if err != nil {
		log.Fatalf("Failed to insert sample data: %v", err)
		return nil
	}

	return db
}

func (suite *RepositoryTestSuite) SetupTest() {
	suite.Context = context.Background()
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	mockDB := NewMockDatabase()
	suite.repository = NewSQLiteRepository(mockDB)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) TestCreate() {
	mockDevice := domain.NewMockDevice()
	err := suite.repository.Save(mockDevice)
	suite.Nil(err)
}
func (suite *RepositoryTestSuite) TestFind() {
	mockDevice := domain.NewMockDevice()
	err := suite.repository.Save(mockDevice)
	suite.Nil(err)
	device, err := suite.repository.FindById(1)
	suite.Nil(err)
	suite.NotNil(device)
}
func (suite *RepositoryTestSuite) TestUpdate() {
	mockDevice := domain.NewMockDevice()
	err := suite.repository.Save(mockDevice)
	suite.Nil(err)
	mockDevice.Name = "teste"
	err = suite.repository.Update(mockDevice)
	suite.Nil(err)
}
func (suite *RepositoryTestSuite) TestDelete() {
	mockDevice := domain.NewMockDevice()
	err := suite.repository.Save(mockDevice)
	suite.Nil(err)
	err = suite.repository.Delete(mockDevice.ID)
	suite.Nil(err)
}

func (suite *RepositoryTestSuite) TestFindById() {
	mockDevice := domain.NewMockDevice()
	err := suite.repository.Save(mockDevice)
	suite.Nil(err)
	device, err := suite.repository.FindById(1)
	suite.Nil(err)
	suite.NotNil(device)
}

func (suite *RepositoryTestSuite) TestFindAll() {
	mockDevice := domain.NewMockDevice()
	err := suite.repository.Save(mockDevice)
	suite.Nil(err)
	devices, err := suite.repository.FindAll()
	suite.Nil(err)
	suite.NotNil(devices)
}

func TestNewDatabaseConnection_Success(t *testing.T) {
	filepath := ":memory:"
	db, err := NewDatabaseConnection(filepath)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if db == nil {
		t.Fatal("expected a non-nil database connection")
	}
}
