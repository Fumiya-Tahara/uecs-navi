package converter

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/dto"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
)

func ToDtoEdgeResponse(edges []domain.Edge) *[]dto.EdgeResponse {
	edgesResponse := make([]dto.EdgeResponse, len(edges))
	for i, edge := range edges {
		edgesResponse[i] = *dto.NewEdgeResponse(edge.ID, edge.WorkflowID, edge.SourceNodeID, edge.TargetNodeID)
	}

	return &edgesResponse
}
