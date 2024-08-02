package usecase

import (
	"context"
	"time"
	"webserver/src/app/domain"
	"webserver/src/app/port"
	"webserver/src/app/repository"
)

type DeviceHandler struct {
	dbRepository *repository.DatabaseRepository
}

func NewDeviceHandler(dbRepository *repository.DatabaseRepository) *DeviceHandler {
	return &DeviceHandler{
		dbRepository: dbRepository,
	}
}

func (d DeviceHandler) InitializeRepository() error {
	err := d.dbRepository.Initialize()
	if err != nil {
		return err
	}
	return nil
}

func (d DeviceHandler) AddDevice(ctx context.Context, device domain.Device) error {
	creationDate := time.Now()
	device.CreationTime = creationDate
	err := d.dbRepository.Save(device)
	if err != nil {
		return err
	}
	return nil
}

func (d DeviceHandler) GetDeviceByID(ctx context.Context, ID int) (domain.Device, error) {
	device, err := d.dbRepository.FindById(ID)
	if err != nil {
		return domain.Device{}, err
	}
	return device, nil
}

func (d DeviceHandler) ListDevices(ctx context.Context) ([]domain.Device, error) {
	devices, err := d.dbRepository.FindAll()
	if err != nil {
		return []domain.Device{}, err
	}
	return devices, nil
}

func (d DeviceHandler) UpdateDevice(ctx context.Context, id int, device domain.Device) error {
	/*deviceInfo, err := d.dbRepository.FindById(id)
	if err != nil {
		return err
	}
	*/
	err := d.dbRepository.Update(device)
	if err != nil {
		return err
	}
	return nil
}

func (d DeviceHandler) PartialUpdateDevice(ctx context.Context, id int, updates map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (d DeviceHandler) DeleteDevice(ctx context.Context, id int) error {
	err := d.dbRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (d DeviceHandler) SearchDevicesByBrand(ctx context.Context, brand string) ([]domain.Device, error) {
	devices, err := d.dbRepository.FindByBrand(brand)
	if err != nil {
		return []domain.Device{}, err
	}
	return devices, nil
}

func NewUseCase(dbRepository *repository.DatabaseRepository) port.UseCase {
	return &DeviceHandler{
		dbRepository: dbRepository,
	}
}
