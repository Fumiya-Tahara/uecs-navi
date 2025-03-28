package dto

type NodeResponse struct {
	ID             int                    `json:"id"`
	WorkflowID     int                    `json:"workflow_id"`
	WorkflowNodeID string                 `json:"workflow_node_id"`
	Type           string                 `json:"type"`
	Data           map[string]interface{} `json:"data"`
	PositionX      float32                `json:"position_x"`
	PositionY      float32                `json:"position_y"`
}

func NewNodeResponse(id, workflowID int, workflowNodeID, nodeType string, data map[string]interface{}, positionX, positionY float32) *NodeResponse {
	return &NodeResponse{
		ID:             id,
		WorkflowID:     workflowID,
		WorkflowNodeID: workflowNodeID,
		Type:           nodeType,
		Data:           data,
		PositionX:      positionX,
		PositionY:      positionY,
	}
}
