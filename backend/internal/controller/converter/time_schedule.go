package converter

import (
	"log"
	"time"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/dto"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
)

func ToDtoTimeScheduleResponse(timeSchedule domain.TimeSchedule, workflows []domain.Workflow) *dto.TimeScheduleResponse {
	timeScheduleRows := timeSchedule.Rows
	dtoTimeScheduleRows := make([]dto.TimeScheduleRow, len(timeScheduleRows))
	for i, timeScheduleRow := range timeScheduleRows {
		dtoCondition := ToDtoCondition(timeScheduleRow.Condition)

		dtoTimeScheduleRows[i] = *dto.NewTimeScheduleRow(timeScheduleRow.StartTime.Format("15:04"), timeScheduleRow.EndTime.Format("15:04"), timeScheduleRow.WorkflowID, *dtoCondition)
	}

	dtoWorkflows := make([]dto.Workflow, len(workflows))
	for i, workflow := range workflows {
		dtoWorkflows[i] = *ToDtoWorkflow(workflow)
	}

	timeScheduleRes := dto.NewTimeScheduleResponse(
		timeSchedule.ID,
		timeSchedule.M304ID,
		dtoWorkflows,
		dtoTimeScheduleRows,
	)

	return timeScheduleRes
}

func ToDomainTimeSchedule(timeScheduleReq dto.TimeScheduleRequest) *domain.TimeSchedule {
	timeScheduleRowsReq := timeScheduleReq.TimeScheduleRows
	timeScheduleRows := make([]domain.TimeScheduleRow, len(timeScheduleRowsReq))
	for i, timeScheduleRowReq := range timeScheduleRowsReq {
		undefinedTimeScheduleID := 0
		condition := ToDomainCondition(timeScheduleRowReq.Condition)

		startTime, err := time.Parse("15:04", timeScheduleRowReq.StartTime)
		if err != nil {
			log.Println("Error parsing time:", err)
			return nil
		}

		endTime, err := time.Parse("15:04", timeScheduleRowReq.EndTime)
		if err != nil {
			log.Println("Error parsing time:", err)
			return nil
		}

		timeScheduleRows[i] = *domain.NewTimeScheduleRow(
			undefinedTimeScheduleID,
			startTime,
			endTime,
			timeScheduleRowReq.SelectedWorkflowID,
			*condition,
		)
	}

	timeSchedule := domain.NewTimeSchedule(
		timeScheduleReq.M304ID,
		timeScheduleRows,
	)

	return timeSchedule
}
