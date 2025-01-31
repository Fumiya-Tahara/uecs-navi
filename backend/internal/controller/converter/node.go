package converter

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/dto"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
)

func ToDtoNodeResponse(nodes []domain.Node) *[]dto.NodeResponse {
	nodesRes := make([]dto.NodeResponse, len(nodes))
	for i, node := range nodes {
		nodesRes[i] = *dto.NewNodeResponse(node.ID, node.WorkflowID, node.NodeID, node.Type, node.Data, float32(node.Position.X), float32(node.Position.Y))
	}

	return &nodesRes
}
