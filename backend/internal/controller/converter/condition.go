package converter

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/dto"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
)

func ToDtoCondition(condition domain.Condition) *dto.Condition {
	dtoCondition := dto.NewCondition(
		condition.ClimateDataID,
		condition.ComparisonOperatorID,
		condition.SetPoint,
	)

	return dtoCondition
}

// The fuction for request
func ToDomainCondition(dtoCondition dto.Condition) *domain.Condition {
	undefinedTimeScheduleID := 0
	condition := domain.NewCondition(
		dtoCondition.SelectedClimateDataID,
		undefinedTimeScheduleID,
		dtoCondition.SelectedComparisonOperatorID,
		dtoCondition.SetPoint,
	)

	return condition
}
