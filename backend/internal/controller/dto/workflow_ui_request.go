package dto

type WorkflowUIRequest struct {
	Nodes []NodeRequest `json:"nodes"`
	Edges []EdgeRequest `json:"edges"`
}

func NewWorkflowUIRequest(nodes []NodeRequest, edges []EdgeRequest) *WorkflowUIRequest {
	return &WorkflowUIRequest{
		Nodes: nodes,
		Edges: edges,
	}
}
