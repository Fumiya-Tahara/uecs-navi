package dto

type TimeScheduleResponse struct {
	M304ID        int            `json:"m304_id"`
	Workflows     []Workflow     `json:"Workflows"`
	TimeSchedules []TimeSchedule `json:"time_schedules"`
}

func NewTimeScheduleResponse(m304ID int, workflows []Workflow, timeSchedules []TimeSchedule) *TimeScheduleResponse {
	return &TimeScheduleResponse{
		M304ID:        m304ID,
		Workflows:     workflows,
		TimeSchedules: timeSchedules,
	}
}
