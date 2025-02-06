package converter

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/dto"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
)

func ToDtoTimeScheduleResponse(timeSchedule domain.TimeSchedule) *dto.TimeScheduleResponse {
	timeScheduleRows := timeSchedule.Rows
	dtoTimeScheduleRows := make([]dto.TimeScheduleRow, len(timeScheduleRows))
	for i, timeScheduleRow := range timeScheduleRows {

		condition := dto.NewCondition(
			timeScheduleRow.Condition.ClimateDataID,
			timeScheduleRow.Condition.ComparisonOperatorID,
			timeScheduleRow.Condition.SetPoint,
		)

		dtoTimeScheduleRows[i] = *dto.NewTimeScheduleRow(timeScheduleRow.StartTime.Format("15:04"), timeScheduleRow.EndTime.Format("15:04"), timeScheduleRow.WorkflowID, *condition)
	}

	timeScheduleRes := dto.NewTimeScheduleResponse(
		timeSchedule.ID,
		timeSchedule.M304ID,
		dtoTimeScheduleRows,
	)

	return timeScheduleRes
}
