package dto

type ClimateData struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Unit string `unit:"unit"`
}

func NewClimateData(id int, name, unit string) *ClimateData {
	return &ClimateData{
		ID:   id,
		Name: name,
		Unit: unit,
	}
}
