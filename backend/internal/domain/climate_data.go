package domain

type ClimateData struct {
	ID   int
	Name string
	Unit string
}

func NewClimateData(id int, name, unit string) *ClimateData {
	return &ClimateData{
		ID:   id,
		Name: name,
		Unit: unit,
	}
}
