package domain

import "time"

type TimeScheduleRow struct {
	ID             int
	TimeScheduleID int
	StartTime      time.Time
	EndTime        time.Time
	WorkflowID     int
}

type TimeSchedule struct {
	ID     int
	M304ID int
	Rows   []TimeScheduleRow
}

func NewTimeScheduleRowWithID(id, timeScheduleID int, startTime time.Time, endTime time.Time, workflowID int) *TimeScheduleRow {
	return &TimeScheduleRow{
		ID:             id,
		TimeScheduleID: timeScheduleID,
		StartTime:      startTime,
		EndTime:        endTime,
		WorkflowID:     workflowID,
	}
}

func NewTimeScheduleRow(timeScheduleID int, startTime time.Time, endTime time.Time, workflowID int) *TimeScheduleRow {
	return &TimeScheduleRow{
		TimeScheduleID: timeScheduleID,
		StartTime:      startTime,
		EndTime:        endTime,
		WorkflowID:     workflowID,
	}
}

func NewTimeScheduleWithID(id, m304ID int, rows []TimeScheduleRow) *TimeSchedule {
	return &TimeSchedule{
		ID:     id,
		M304ID: m304ID,
		Rows:   rows,
	}
}

func NewTimeSchedule(m304ID int, rows []TimeScheduleRow) *TimeSchedule {
	return &TimeSchedule{
		M304ID: m304ID,
		Rows:   rows,
	}
}
