package domain

type Condition struct {
	ID                   int
	ClimateDataID        int
	TimeScheduleRowID    int
	ComparisonOperatorID int
	SetPoint             float32
}

func NewConditionWithID(id, climateDataID, timeScheduleID, comparisonOperatorID int, setPoint float32) *Condition {
	return &Condition{
		ID:                   id,
		ClimateDataID:        climateDataID,
		TimeScheduleRowID:    timeScheduleID,
		ComparisonOperatorID: comparisonOperatorID,
		SetPoint:             setPoint,
	}
}

func NewCondition(climateDataID, timeScheduleID, comparisonOperatorID int, setPoint float32) *Condition {
	return &Condition{
		ClimateDataID:        climateDataID,
		TimeScheduleRowID:    timeScheduleID,
		ComparisonOperatorID: comparisonOperatorID,
		SetPoint:             setPoint,
	}
}
