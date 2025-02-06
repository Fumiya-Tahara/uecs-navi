package dto

type TimeScheduleResponse struct {
	ID               int
	M304ID           int               `json:"m304_id"`
	TimeScheduleRows []TimeScheduleRow `json:"time_schedules"`
}

func NewTimeScheduleResponse(id, m304ID int, timeSchedules []TimeScheduleRow) *TimeScheduleResponse {
	return &TimeScheduleResponse{
		ID:               id,
		M304ID:           m304ID,
		TimeScheduleRows: timeSchedules,
	}
}
