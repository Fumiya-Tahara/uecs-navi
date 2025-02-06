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

func ToDomainEdge(edgesReq []dto.EdgeRequest) *[]domain.Edge {
	edges := make([]domain.Edge, len(edgesReq))
	for i, edgeReq := range edgesReq {
		undefinedWorkflowID := 0

		edges[i] = *domain.NewEdge(undefinedWorkflowID, edgeReq.SourceNodeID, edgeReq.TargetNodeID)
	}

	return &edges
}
