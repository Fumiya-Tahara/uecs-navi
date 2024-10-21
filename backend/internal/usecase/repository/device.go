package repository

import (
	"context"
	"database/sql"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type JoinedDevice struct {
	ID          int
	HouseID     int
	DeviceName  *string
	SetPoint    *float64
	Duration    *int
	ClimateData string
	Unit        string
}

func NewJoinedDevice(id int, houseID int, deviceName *string, setPoint *float64, duration *int, climateData, unit string) *JoinedDevice {
	return &JoinedDevice{
		ID:          id,
		HouseID:     houseID,
		DeviceName:  deviceName,
		SetPoint:    setPoint,
		Duration:    duration,
		ClimateData: climateData,
		Unit:        unit,
	}
}

type DeviceRepository struct {
	queries *mysqlc.Queries
}

func NewDeviceRepository(queries *mysqlc.Queries) *DeviceRepository {
	return &DeviceRepository{
		queries: queries,
	}
}

func (dr DeviceRepository) CreateDevice(newDevice domain.Device) (int64, error) {
	ctx := context.Background()

	arg := mysqlc.CreateDeviceParams{
		HouseID:       int32(newDevice.HouseID),
		ClimateDataID: int32(newDevice.ClimateDataID),
		DeviceName: sql.NullString{
			String: *newDevice.DeviceName,
			Valid:  newDevice.DeviceName != nil,
		},
		SetPoint: sql.NullFloat64{
			Float64: *newDevice.SetPoint,
			Valid:   newDevice.SetPoint != nil,
		},
		Duration: sql.NullInt32{
			Int32: int32(*newDevice.Duration),
			Valid: newDevice.Duration != nil,
		},
	}

	id, err := dr.queries.CreateDevice(ctx, arg)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (dr DeviceRepository) GetDevicesFromHouse(houseID int) ([]*domain.Device, error) {
	ctx := context.Background()

	devicesRow, err := dr.queries.GetDevicesFromHouse(ctx, int32(houseID))
	if err != nil {
		return nil, err
	}

	devices := make([]*domain.Device, len(devicesRow))
	for i, v := range devicesRow {
		var deviceName *string
		if v.DeviceName.Valid {
			deviceName = &v.DeviceName.String
		}
		var setPoint *float64
		if v.SetPoint.Valid {
			setPoint = &v.SetPoint.Float64
		}
		var duration *int
		if v.Duration.Valid {
			temp := int(v.Duration.Int32)
			duration = &temp
		}
		devices[i] = domain.NewDeviceWithID(
			int(v.ID),
			int(v.HouseID),
			int(v.ClimateDataID),
			deviceName,
			setPoint,
			duration,
		)
	}

	return devices, nil
}

func (dr DeviceRepository) GetJoinedDevicesFromHouse(houseID int) ([]*JoinedDevice, error) {
	ctx := context.Background()

	joinedDevicesRow, err := dr.queries.GetJoinedDevicesFromHouse(ctx, int32(houseID))
	if err != nil {
		return nil, err
	}

	joinedDevices := make([]*JoinedDevice, len(joinedDevicesRow))
	for i, v := range joinedDevicesRow {
		var deviceName *string
		if v.DeviceName.Valid {
			deviceName = &v.DeviceName.String
		}
		var setPoint *float64
		if v.SetPoint.Valid {
			setPoint = &v.SetPoint.Float64
		}
		var duration *int
		if v.Duration.Valid {
			temp := int(v.Duration.Int32)
			duration = &temp
		}
		joinedDevices[i] = NewJoinedDevice(
			int(v.ID),
			int(v.HouseID),
			deviceName,
			setPoint,
			duration,
			v.ClimateDataName,
			v.Unit,
		)
	}

	return joinedDevices, nil
}
