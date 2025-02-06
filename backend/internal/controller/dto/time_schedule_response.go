package dto

type TimeScheduleResponse struct {
	ID               int               `json:"id"`
	M304ID           int               `json:"m304_id"`
	Workflows        []Workflow        `json:"workflows"`
	TimeScheduleRows []TimeScheduleRow `json:"time_schedules"`
}

func NewTimeScheduleResponse(id, m304ID int, workflows []Workflow, timeSchedules []TimeScheduleRow) *TimeScheduleResponse {
	return &TimeScheduleResponse{
		ID:               id,
		M304ID:           m304ID,
		Workflows:        workflows,
		TimeScheduleRows: timeSchedules,
	}
}
