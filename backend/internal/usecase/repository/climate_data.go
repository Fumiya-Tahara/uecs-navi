package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type ClimateDataRepository struct {
	queries *mysqlc.Queries
}

func NewClimateDataRepository(queries *mysqlc.Queries) *ClimateDataRepository {
	return &ClimateDataRepository{
		queries: queries,
	}
}

func (cdr ClimateDataRepository) GetAllClimateData() (*[]domain.ClimateData, error) {
	ctx := context.Background()

	climateDataRows, err := cdr.queries.GetAllClimateData(ctx)
	if err != nil {
		return nil, err
	}

	climateData := make([]domain.ClimateData, len(climateDataRows))
	for i, v := range climateDataRows {
		climateData[i] = *domain.NewClimateData(
			int(v.ID),
			v.Name,
			v.Unit,
		)
	}

	return &climateData, nil
}

func (cdr ClimateDataRepository) GetClimateDataFromID(id int) (*domain.ClimateData, error) {
	ctx := context.Background()

	climateData, err := cdr.queries.GetClimateDataFromID(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	getClimateData := domain.NewClimateData(
		int(climateData.ID),
		climateData.Name,
		climateData.Unit,
	)

	return getClimateData, nil
}
