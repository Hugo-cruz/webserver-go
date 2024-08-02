package port

import "webserver/src/app/domain"
import "context"

type UseCase interface {
	InitializeRepository() error
	AddDevice(ctx context.Context, device domain.Device) error
	GetDeviceByID(ctx context.Context, ID int) (domain.Device, error)
	ListDevices(ctx context.Context) ([]domain.Device, error)
	UpdateDevice(ctx context.Context, id int, device domain.Device) error
	PartialUpdateDevice(ctx context.Context, id int, updates map[string]interface{}) error
	DeleteDevice(ctx context.Context, id int) error
	SearchDevicesByBrand(ctx context.Context, brand string) ([]domain.Device, error)
}
