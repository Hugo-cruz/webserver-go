package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
	"webserver/src/app/domain"
	"webserver/src/app/port/device"
	repository "webserver/src/app/repository/mock"
)

type DeviceUseCaseTestSuite struct {
	suite.Suite
	context.Context
	UseCase    device.UseCase
	repository *repository.MockRepository
}

func (suite *DeviceUseCaseTestSuite) SetupTest() {
	suite.Context = context.Background()
	ctrl := gomock.NewController(suite.T())
	defer ctrl.Finish()
	suite.repository = repository.NewMockRepository(ctrl)
	suite.UseCase = NewUseCase(suite.repository)
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(DeviceUseCaseTestSuite))
}

func (suite *DeviceUseCaseTestSuite) TestGetDeviceByID() {
	mockDevice := domain.NewMockDevice()
	suite.repository.EXPECT().FindById(1).Return(mockDevice, nil)
	deviceResult, err := suite.UseCase.GetDeviceByID(suite.Context, 1)
	suite.Nil(err)
	suite.Equal(mockDevice, deviceResult)
}
func (suite *DeviceUseCaseTestSuite) TestGetDeviceByIDNotFound() {
	suite.repository.EXPECT().FindById(0).Return(&domain.Device{}, errors.New("not found"))
	_, err := suite.UseCase.GetDeviceByID(suite.Context, 0)
	suite.NotNil(err)
}
func (suite *DeviceUseCaseTestSuite) TestGetDeviceByIDError() {
	suite.repository.EXPECT().FindById(0).Return(&domain.Device{}, errors.New("error"))
	_, err := suite.UseCase.GetDeviceByID(suite.Context, 0)
	suite.NotNil(err)
}

func (suite *DeviceUseCaseTestSuite) TestDeleteDevice() {
	suite.repository.EXPECT().Delete(1).Return(nil)
	err := suite.UseCase.DeleteDevice(suite.Context, 1)
	suite.Nil(err)
}
func (suite *DeviceUseCaseTestSuite) TestDeleteDeviceError() {
	suite.repository.EXPECT().Delete(1).Return(errors.New("error"))
	err := suite.UseCase.DeleteDevice(suite.Context, 1)
	suite.NotNil(err)
}
func (suite *DeviceUseCaseTestSuite) TestGetDeviceList() {
	mockDevice := domain.NewMockDevice()
	suite.repository.EXPECT().FindAll().Return([]*domain.Device{mockDevice}, nil)
	devices, err := suite.UseCase.ListDevices(suite.Context)
	suite.Nil(err)
	suite.Equal(mockDevice, devices[0])
}
func (suite *DeviceUseCaseTestSuite) TestGetDeviceListError() {
	suite.repository.EXPECT().FindAll().Return([]*domain.Device{}, errors.New("error"))
	_, err := suite.UseCase.ListDevices(suite.Context)
	suite.NotNil(err)
}
func (suite *DeviceUseCaseTestSuite) TestUpdateDevice() {
	mockDevice := domain.NewMockDevice()
	suite.repository.EXPECT().Update(mockDevice).Return(nil)
	err := suite.UseCase.UpdateDevice(suite.Context, 1, mockDevice)
	suite.Nil(err)
}
func (suite *DeviceUseCaseTestSuite) TestUpdateDeviceError() {
	mockDevice := domain.NewMockDevice()
	suite.repository.EXPECT().Update(mockDevice).Return(errors.New("error"))
	err := suite.UseCase.UpdateDevice(suite.Context, 1, mockDevice)
	suite.NotNil(err)
}
