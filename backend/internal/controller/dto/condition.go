package dto

type Condition struct {
	SelectedClimateDataID        int     `json:"selected_climate_data_id"`
	SelectedComparisonOperatorID int     `json:"selected_comparison_operator_id"`
	SetPoint                     float32 `json:"set_point"`
}

func NewCondition(selectedClimateDataID, selectedComparisonOperatorID int, setPoint float32) *Condition {
	return &Condition{
		SelectedClimateDataID:        selectedClimateDataID,
		SelectedComparisonOperatorID: selectedComparisonOperatorID,
		SetPoint:                     setPoint,
	}
}
