package dto

type EdgeRequest struct {
	SourceNodeID string `json:"source_node_id"`
	TargetNodeID string `json:"target_node_id"`
}

func NewEdgeRequest(sourceNodeID, targetNodeID string) *EdgeRequest {
	return &EdgeRequest{
		SourceNodeID: sourceNodeID,
		TargetNodeID: targetNodeID,
	}
}
