package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type ConditionRepository struct {
	queries *mysqlc.Queries
}

func NewConditionRepository(queries *mysqlc.Queries) *ConditionRepository {
	return &ConditionRepository{
		queries: queries,
	}
}

func (cr ConditionRepository) CreateCondition(newCondition *domain.Condition) (int, error) {
	ctx := context.Background()

	arg := mysqlc.CreateConditionParams{
		ClimateDataID:        int32(newCondition.ClimateDataID),
		TimeScheduleRowID:    int32(newCondition.TimeScheduleRowID),
		ComparisonOperatorID: int32(newCondition.ComparisonOperatorID),
		SetPoint:             float64(newCondition.SetPoint),
	}

	id, err := cr.queries.CreateCondition(ctx, arg)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (cr ConditionRepository) GetConditionFromTimeScheduleRow(timeScheduleRowID int) (*domain.Condition, error) {
	ctx := context.Background()

	conditionRow, err := cr.queries.GetConditionFromTimeScheduleRow(ctx, int32(timeScheduleRowID))
	if err != nil {
		return nil, err
	}

	condition := domain.NewConditionWithID(
		int(conditionRow.ID),
		int(conditionRow.ClimateDataID),
		int(conditionRow.TimeScheduleRowID),
		int(conditionRow.ComparisonOperatorID),
		float32(conditionRow.SetPoint),
	)

	return condition, nil
}

func (cr ConditionRepository) UpdateCondition(condition domain.Condition) error {
	ctx := context.Background()

	arg := mysqlc.UpdateConditionParams{
		ClimateDataID:        int32(condition.ClimateDataID),
		TimeScheduleRowID:    int32(condition.TimeScheduleRowID),
		ComparisonOperatorID: int32(condition.ComparisonOperatorID),
		SetPoint:             float64(condition.SetPoint),
		ID:                   int32(condition.ID),
	}

	if err := cr.queries.UpdateCondition(ctx, arg); err != nil {
		return err
	}

	return nil
}
