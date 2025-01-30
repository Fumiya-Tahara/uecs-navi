package dto

type TimeSchedule struct {
	StartTime          string    `json:"start_time"`
	EndTime            string    `json:"end_time"`
	SelectedWorkflowID int       `json:"selected_workflow_id"`
	Condition          Condition `json:"condition"`
}

func NewTimeSchedule(startTime, endTime string, selectedWorkflowID int, condition Condition) *TimeSchedule {
	return &TimeSchedule{
		StartTime:          startTime,
		EndTime:            endTime,
		SelectedWorkflowID: selectedWorkflowID,
		Condition:          condition,
	}
}
