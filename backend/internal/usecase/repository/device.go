package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/utils"
)

type DeviceRepository struct {
	queries *mysqlc.Queries
}

func NewDeviceRepository(queries *mysqlc.Queries) *DeviceRepository {
	return &DeviceRepository{
		queries: queries,
	}
}

func (dr DeviceRepository) CreateDevice(newDevice domain.Device) (int, error) {
	ctx := context.Background()

	arg := mysqlc.CreateDeviceParams{
		ClimateDataID: int32(newDevice.ClimateDataID),
		M304ID:        int32(newDevice.M304ID),
		Name:          *newDevice.Name,
		Rly:           utils.PointerToNullInt32(newDevice.Rly),
	}

	id, err := dr.queries.CreateDevice(ctx, arg)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (dr DeviceRepository) GetDevicesFromM304(m304ID int) (*[]domain.Device, error) {
	ctx := context.Background()

	deviceRows, err := dr.queries.GetDevicesFromM304(ctx, int32(m304ID))
	if err != nil {
		return nil, err
	}

	devices := make([]domain.Device, len(deviceRows))
	for i, v := range deviceRows {
		rly := utils.NullInt32ToPointer(v.Rly)

		devices[i] = *domain.NewDeviceWithID(
			int(v.ID),
			int(v.ClimateDataID),
			int(v.M304ID),
			&v.Name,
			rly,
		)
	}

	return &devices, nil
}
