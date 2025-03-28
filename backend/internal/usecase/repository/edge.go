package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type EdgeRepository struct {
	queries *mysqlc.Queries
}

func NewEdgeRepository(queries *mysqlc.Queries) *EdgeRepository {
	return &EdgeRepository{
		queries: queries,
	}
}

func (er EdgeRepository) CreateEdge(newEdge domain.Edge) (int, error) {
	ctx := context.Background()

	arg := mysqlc.CreateEdgeParams{
		WorkflowID:   int32(newEdge.WorkflowID),
		SourceNodeID: newEdge.SourceNodeID,
		TargetNodeID: newEdge.TargetNodeID,
	}

	id, err := er.queries.CreateEdge(ctx, arg)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (er EdgeRepository) GetEdgesFromWorkflow(workflowID int) (*[]domain.Edge, error) {
	ctx := context.Background()

	edgeRows, err := er.queries.GetEdgesFromWorkflow(ctx, int32(workflowID))
	if err != nil {
		return nil, err
	}

	edges := make([]domain.Edge, len(edgeRows))
	for i, v := range edgeRows {
		edges[i] = *domain.NewEdgeWithID(
			int(v.ID),
			int(v.WorkflowID),
			v.SourceNodeID,
			v.TargetNodeID,
		)
	}

	return &edges, nil
}

func (er EdgeRepository) UpdateEdge(edge domain.Edge) error {
	ctx := context.Background()

	arg := mysqlc.UpdateEdgeParams{
		WorkflowID:   int32(edge.WorkflowID),
		SourceNodeID: edge.SourceNodeID,
		TargetNodeID: edge.TargetNodeID,
		ID:           int32(edge.ID),
	}

	if err := er.queries.UpdateEdge(ctx, arg); err != nil {
		return err
	}

	return nil
}
