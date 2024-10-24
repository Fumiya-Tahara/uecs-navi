package mocks

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/repository"
)

type MockDeviceRepository struct {
	DeviceTable       map[int][]*domain.Device
	JoinedDeviceTable map[int][]*repository.JoinedDevice
}

func NewMockDeviceRepository() *MockDeviceRepository {
	deviceTableMap := make(map[int][]*domain.Device)
	var deviceName *string
	var setPoint *float64
	var duration *int
	temp_DN1 := "加温器"
	temp_SP1 := float64(22)
	temp_Du1 := 1
	deviceName = &temp_DN1
	setPoint = &temp_SP1
	duration = &temp_Du1
	device1 := &domain.Device{
		ID:            1,
		HouseID:       1,
		ClimateDataID: 1,
		DeviceName:    deviceName,
		SetPoint:      setPoint,
		Duration:      duration,
	}
	temp_DN2 := "加湿器"
	temp_SP2 := float64(60)
	temp_Du2 := 2
	deviceName = &temp_DN2
	setPoint = &temp_SP2
	duration = &temp_Du2
	device2 := &domain.Device{
		ID:            2,
		HouseID:       2,
		ClimateDataID: 2,
		DeviceName:    deviceName,
		SetPoint:      setPoint,
		Duration:      duration,
	}
	temp_DN3 := "CO2濃度センサ"
	temp_SP3 := float64(420)
	temp_Du3 := 3
	deviceName = &temp_DN3
	setPoint = &temp_SP3
	duration = &temp_Du3
	device3 := &domain.Device{
		ID:            3,
		HouseID:       2,
		ClimateDataID: 3,
		DeviceName:    deviceName,
		SetPoint:      setPoint,
		Duration:      duration,
	}
	temp_DN4 := "CO2濃度センサ"
	temp_SP4 := float64(420)
	temp_Du4 := 4
	deviceName = &temp_DN4
	setPoint = &temp_SP4
	duration = &temp_Du4
	device4 := &domain.Device{
		ID:            4,
		HouseID:       3,
		ClimateDataID: 3,
		DeviceName:    deviceName,
		SetPoint:      setPoint,
		Duration:      duration,
	}

	deviceTableMap[1] = []*domain.Device{device1}
	deviceTableMap[2] = []*domain.Device{device2, device3}
	deviceTableMap[3] = []*domain.Device{device4}

	joinedDeviceTableMap := make(map[int][]*repository.JoinedDevice)

	temp_SP5 := float64(22)
	temp_Du5 := 1
	setPoint = &temp_SP5
	duration = &temp_Du5
	joinedDevice1 := &repository.JoinedDevice{
		ID:          1,
		HouseID:     1,
		SetPoint:    setPoint,
		Duration:    duration,
		ClimateData: "気温",
		Unit:        "℃",
	}
	temp_SP6 := float64(60)
	temp_Du6 := 2
	setPoint = &temp_SP6
	duration = &temp_Du6
	joinedDevice2 := &repository.JoinedDevice{
		ID:          2,
		HouseID:     2,
		SetPoint:    setPoint,
		Duration:    duration,
		ClimateData: "湿度",
		Unit:        "%",
	}
	temp_SP7 := float64(420)
	temp_Du7 := 3
	setPoint = &temp_SP7
	duration = &temp_Du7
	joinedDevice3 := &repository.JoinedDevice{
		ID:          3,
		HouseID:     2,
		SetPoint:    setPoint,
		Duration:    duration,
		ClimateData: "二酸化炭素量",
		Unit:        "ppm",
	}
	temp_SP8 := float64(420)
	temp_Du8 := 4
	setPoint = &temp_SP8
	duration = &temp_Du8
	joinedDevice4 := &repository.JoinedDevice{
		ID:          4,
		HouseID:     3,
		SetPoint:    setPoint,
		Duration:    duration,
		ClimateData: "二酸化炭素量",
		Unit:        "ppm",
	}

	joinedDeviceTableMap[1] = []*repository.JoinedDevice{joinedDevice1}
	joinedDeviceTableMap[2] = []*repository.JoinedDevice{joinedDevice2, joinedDevice3}
	joinedDeviceTableMap[3] = []*repository.JoinedDevice{joinedDevice4}

	return &MockDeviceRepository{
		DeviceTable:       deviceTableMap,
		JoinedDeviceTable: joinedDeviceTableMap,
	}

}

func (dr MockDeviceRepository) GetDevicesFromHouse(houseID int) ([]*domain.Device, error) {
	return dr.DeviceTable[houseID], nil
}

func (dr MockDeviceRepository) GetJoinedDevicesFromHouse(houseID int) ([]*repository.JoinedDevice, error) {
	return dr.JoinedDeviceTable[houseID], nil
}

func (dr MockDeviceRepository) CreateDevice(newDevice domain.Device) (int64, error) {
	return 0, nil
}
