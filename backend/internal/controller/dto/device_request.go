package dto

type DeviceRequest struct {
	DeviceName    string  `json:"device_name"`
	ClimateDataID int     `json:"climate_data_id"`
	Unit          string  `json:"unit"`
	SetPoint      float32 `json:"set_point"`
	Duration      int     `json:"duration"`
}

func NewDeviceRequest(deviceName string, climateDataID int, unit string, setPoint float32, duration int) *DeviceRequest {
	return &DeviceRequest{
		DeviceName:    deviceName,
		ClimateDataID: climateDataID,
		Unit:          unit,
		SetPoint:      setPoint,
		Duration:      duration,
	}
}
