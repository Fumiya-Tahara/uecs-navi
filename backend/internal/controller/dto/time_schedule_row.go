package dto

type TimeScheduleRow struct {
	StartTime          string    `json:"start_time"`
	EndTime            string    `json:"end_time"`
	SelectedWorkflowID int       `json:"selected_workflow_id"`
	Condition          Condition `json:"condition"`
}

func NewTimeScheduleRow(startTime, endTime string, selectedWorkflowID int, condition Condition) *TimeScheduleRow {
	return &TimeScheduleRow{
		StartTime:          startTime,
		EndTime:            endTime,
		SelectedWorkflowID: selectedWorkflowID,
		Condition:          condition,
	}
}
