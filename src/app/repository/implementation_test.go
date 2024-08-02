package repository

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/net/context"
	"gorm.io/gorm"
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
	return db
}

func (suite *RepositoryTestSuite) SetupTest() {
	suite.Context = context.Background()
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	mockDB := NewMockDatabase()
	suite.repository = NewDatabaseRepository(mockDB)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}

func (suite *RepositoryTestSuite) TestCreate() {
	mockDevice := NewMockDevice()
	err := suite.repository.Save(mockDevice)
	suite.Nil(err)
}
func (suite *RepositoryTestSuite) TestFind() {
	mockDevice := NewMockDevice()
	err := suite.repository.Save(mockDevice)
	suite.Nil(err)
	device, err := suite.repository.FindById(1)
	suite.Nil(err)
	suite.NotNil(device)
}
func (suite *RepositoryTestSuite) TestUpdate() {
	mockDevice := NewMockDevice()
	err := suite.repository.Save(mockDevice)
	suite.Nil(err)
	mockDevice.Name = "teste"
	err = suite.repository.Update(mockDevice)
	suite.Nil(err)
}
func (suite *RepositoryTestSuite) TestDelete() {
	mockDevice := NewMockDevice()
	err := suite.repository.Save(mockDevice)
	suite.Nil(err)
	err = suite.repository.Delete(mockDevice.ID)
	suite.Nil(err)
}

func (suite *RepositoryTestSuite) TestFindById() {
	mockDevice := NewMockDevice()
	err := suite.repository.Save(mockDevice)
	suite.Nil(err)
	device, err := suite.repository.FindById(1)
	suite.Nil(err)
	suite.NotNil(device)
}

func (suite *RepositoryTestSuite) TestFindAll() {
	mockDevice := NewMockDevice()
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

func NewMockDevice() domain.Device {
	return domain.Device{
		ID:    1,
		Name:  "Teste",
		Brand: "Brand",
	}
}
