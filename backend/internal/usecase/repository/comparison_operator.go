package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type ComparisonOperatorRepository struct {
	queries *mysqlc.Queries
}

func NewComparisonOperatorRepository(queries *mysqlc.Queries) *ComparisonOperatorRepository {
	return &ComparisonOperatorRepository{
		queries: queries,
	}
}

func (cor ComparisonOperatorRepository) GetAllComparisonOperators() (*[]domain.ComparisonOperator, error) {
	ctx := context.Background()

	comparisonOperatorRows, err := cor.queries.GetAllComparisonOperators(ctx)
	if err != nil {
		return nil, err
	}

	comparisonOperators := make([]domain.ComparisonOperator, len(comparisonOperatorRows))
	for i, v := range comparisonOperatorRows {
		comparisonOperators[i] = *domain.NewComparisonOperator(
			int(v.ID),
			v.ComparisonOperator,
		)
	}

	return &comparisonOperators, nil
}

func (cor ComparisonOperatorRepository) GetAllComparisonOperatorFromID(id int) (*domain.ComparisonOperator, error) {
	ctx := context.Background()

	comparisonOperatorRow, err := cor.queries.GetComparisonOperatorFromID(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	comparisonOperator := domain.NewComparisonOperatorWithID(
		int(comparisonOperatorRow.ID),
		comparisonOperatorRow.ComparisonOperator,
	)

	return comparisonOperator, nil
}
