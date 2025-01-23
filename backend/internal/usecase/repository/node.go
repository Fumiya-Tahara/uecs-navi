package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type NodeRepository struct {
	queries *mysqlc.Queries
}

func NewNodeRepository(queries *mysqlc.Queries) *NodeRepository {
	return &NodeRepository{
		queries: queries,
	}
}

func (nr *NodeRepository) CreateNode(newNode domain.Node) (int, error) {
	ctx := context.Background()

	jsonData, err := MapToRawMessage(newNode.Data)
	if err != nil {
		return 0, err
	}

	arg := mysqlc.CreateNodeParams{
		WorkflowID:     int32(newNode.WorkflowID),
		WorkflowNodeID: newNode.NodeID,
		Type:           newNode.Type,
		Data:           jsonData,
	}

	id, err := nr.queries.CreateNode(ctx, arg)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (nr *NodeRepository) GetNodesFromWorkflow(workflowID int) ([]*domain.Node, error) {
	ctx := context.Background()

	nodeRows, err := nr.queries.GetNodesFromWorkflow(ctx, int32(workflowID))
	if err != nil {
		return nil, err
	}

	nodes := make([]*domain.Node, len(nodeRows))
	for i, v := range nodeRows {
		data, err := RawMessageToMap(v.Data)
		if err != nil {
			continue
		}

		position := domain.NewPosition(v.PositionX, v.PositionY)

		nodes[i] = domain.NewNode(
			int(v.WorkflowID),
			v.WorkflowNodeID,
			v.Type,
			data,
			*position,
		)
	}

	return nodes, nil
}
