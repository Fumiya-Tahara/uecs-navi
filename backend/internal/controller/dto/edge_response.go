package dto

type EdgeResponse struct {
	ID           int    `json:"id"`
	WorkflowID   int    `json:"workflow_id"`
	SourceNodeID string `json:"source_node_id"`
	TargetNodeID string `json:"target_node_id"`
}

func NewEdgeResponse(id, workflowID int, sourceNodeID, targetNodeID string) *EdgeResponse {
	return &EdgeResponse{
		ID:           id,
		WorkflowID:   workflowID,
		SourceNodeID: sourceNodeID,
		TargetNodeID: targetNodeID,
	}
}
