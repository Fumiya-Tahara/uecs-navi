package domain

type Device struct {
	ID            int
	ClimateDataID int
	M304ID        int
	Name          *string
	Rly           *int
}

func NewDeviceWithID(id int, climateDataID int, m304ID int, name *string, rly *int) *Device {
	return &Device{
		ID:            id,
		ClimateDataID: climateDataID,
		M304ID:        m304ID,
		Name:          name,
		Rly:           rly,
	}
}

func NewDevice(climateDataID int, m304ID int, name *string, rly *int) *Device {
	return &Device{
		ClimateDataID: climateDataID,
		M304ID:        m304ID,
		Name:          name,
		Rly:           rly,
	}
}
