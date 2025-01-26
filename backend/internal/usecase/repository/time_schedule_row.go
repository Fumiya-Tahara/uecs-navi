package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type TimeScheduleRowRepository struct {
	queries *mysqlc.Queries
}

func NewTimeScheduleRowRepository(queries *mysqlc.Queries) *TimeScheduleRowRepository {
	return &TimeScheduleRowRepository{
		queries: queries,
	}
}

func (tsrr *TimeScheduleRowRepository) CreateTimeScheduleRow(newTimeScheduleRow domain.TimeScheduleRow) (int, error) {
	ctx := context.Background()

	arg := mysqlc.CreateTimeScheduleRowParams{
		TimeScheduleID: int32(newTimeScheduleRow.TimeScheduleID),
		StartTime:      newTimeScheduleRow.StartTime,
		EndTime:        newTimeScheduleRow.EndTime,
		WorkflowID:     int32(newTimeScheduleRow.WorkflowID),
	}

	id, err := tsrr.queries.CreateTimeScheduleRow(ctx, arg)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (tsrr *TimeScheduleRowRepository) GetTimeScheduleRowsFromTimeSchedule(timeScheduleID int) (*[]domain.TimeScheduleRow, error) {
	ctx := context.Background()

	timeScheduleRowRows, err := tsrr.queries.GetTimeScheduleRowsFromTimeSchedule(ctx, int32(timeScheduleID))
	if err != nil {
		return nil, err
	}

	timeScheduleRows := make([]domain.TimeScheduleRow, len(timeScheduleRowRows))
	for i, v := range timeScheduleRowRows {
		timeScheduleRows[i] = *domain.NewTimeScheduleRowWithID(
			int(v.ID),
			int(v.TimeScheduleID),
			v.StartTime,
			v.EndTime,
			int(v.WorkflowID),
		)
	}

	return &timeScheduleRows, nil
}
