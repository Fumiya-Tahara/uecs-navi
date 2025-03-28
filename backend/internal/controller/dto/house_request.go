package dto

type HouseRequest string

func NewHouseRequest(name string) *HouseRequest {
	request := HouseRequest(name)
	return &request
}
