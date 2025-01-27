package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type TimeScheduleRepository struct {
	queries *mysqlc.Queries
}

func NewTimeScheduleRepository(queries *mysqlc.Queries) *TimeScheduleRepository {
	return &TimeScheduleRepository{
		queries: queries,
	}
}

func (tsr TimeScheduleRepository) CreateTimeSchedule(m304ID int) (int, error) {
	ctx := context.Background()

	id, err := tsr.queries.CreateTimeSchedule(ctx, int32(m304ID))
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (tsr TimeScheduleRepository) GetTimeScheduleFromM304(m304ID int) (*domain.TimeSchedule, error) {
	ctx := context.Background()

	timeScheduleRow, err := tsr.queries.GetTimeScheduleFromM304(ctx, int32(m304ID))
	if err != nil {
		return nil, err
	}

	emptyRows := []domain.TimeScheduleRow{}

	timeSchedule := domain.NewTimeSchedule(
		int(timeScheduleRow.M304ID),
		emptyRows,
	)

	return timeSchedule, nil
}
