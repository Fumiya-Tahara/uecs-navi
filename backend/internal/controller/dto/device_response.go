package dto

type DeviceResponse struct {
	ID          int     `json:"id"`
	DeviceName  string  `json:"device_name"`
	ClimateData string  `json:"climate_data"`
	Unit        string  `json:"unit"`
	SetPoint    float64 `json:"set_point"`
	Duration    int     `json:"duration"`
}

func NewDeviceResponse(id int, deviceName, climateData, unit string, setPoint float64, duration int) *DeviceResponse {
	return &DeviceResponse{
		ID:          id,
		DeviceName:  deviceName,
		ClimateData: climateData,
		Unit:        unit,
		SetPoint:    setPoint,
		Duration:    duration,
	}
}
