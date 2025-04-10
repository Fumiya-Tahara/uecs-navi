package testdata

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/utils"
)

func GetDevicesMap() map[int][]domain.Device {
	deviceMap := map[int][]domain.Device{
		1: {
			*domain.NewDeviceWithID(1, 1, 1, utils.StringPtr("Device1"), utils.IntPtr(1)),
			*domain.NewDeviceWithID(2, 2, 1, utils.StringPtr("Device2"), utils.IntPtr(0)),
		},
		2: {
			*domain.NewDeviceWithID(3, 1, 2, utils.StringPtr("Device3"), utils.IntPtr(1)),
			*domain.NewDeviceWithID(4, 2, 2, utils.StringPtr("Device4"), utils.IntPtr(0)),
		},
		3: {
			*domain.NewDeviceWithID(5, 1, 3, utils.StringPtr("Device5"), utils.IntPtr(1)),
		},
	}

	return deviceMap
}
