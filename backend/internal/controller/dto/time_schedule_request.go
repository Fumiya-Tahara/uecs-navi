package dto

type TimeScheduleRequest struct {
	M304ID       int            `json:"m304_id"`
	TimeSchedule []TimeSchedule `json:"time_schedule"`
}

func NewTimeScheduleRequest(m304ID int, timeSchedules []TimeSchedule) *TimeScheduleRequest {
	return &TimeScheduleRequest{
		M304ID:       m304ID,
		TimeSchedule: timeSchedules,
	}
}
