package dto

type NodeRequest struct {
	WorkflowNodeID string                 `json:"workflow_node_id"`
	Type           string                 `json:"type"`
	Data           map[string]interface{} `json:"data"`
	PositionX      float32                `json:"position_x"`
	PositionY      float32                `json:"position_y"`
}

func NewNodeRequest(workflowNodeID, nodeType string, data map[string]interface{}, positionX, positionY float32) *NodeRequest {
	return &NodeRequest{
		WorkflowNodeID: workflowNodeID,
		Type:           nodeType,
		Data:           data,
		PositionX:      positionX,
		PositionY:      positionY,
	}
}
