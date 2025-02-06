package dto

type TimeScheduleRequest struct {
	M304ID           int               `json:"m304_id"`
	TimeScheduleRows []TimeScheduleRow `json:"time_schedule"`
}

func NewTimeScheduleRequest(m304ID int, timeSchedules []TimeScheduleRow) *TimeScheduleRequest {
	return &TimeScheduleRequest{
		M304ID:           m304ID,
		TimeScheduleRows: timeSchedules,
	}
}
