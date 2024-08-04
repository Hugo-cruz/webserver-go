package device

import "webserver/src/app/domain"
import "context"

//go:generate mockgen -source=./usecase.go -destination=./mock/usecase_mock.go -package=device
type UseCase interface {
	InitializeRepository() error
	AddDevice(ctx context.Context, device *domain.Device) error
	GetDeviceByID(ctx context.Context, ID int) (*domain.Device, error)
	ListDevices(ctx context.Context) ([]*domain.Device, error)
	UpdateDevice(ctx context.Context, id int, device *domain.Device) error
	DeleteDevice(ctx context.Context, id int) error
	SearchDevicesByBrand(ctx context.Context, brand string) ([]domain.Device, error)
}
