package dto

type HousesResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewHousesResponse(id int, name string) *HousesResponse {
	return &HousesResponse{
		ID:   id,
		Name: name,
	}
}
