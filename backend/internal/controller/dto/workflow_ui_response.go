package dto

type WorkflowUIResponse struct {
	Nodes []NodeResponse `json:"nodes"`
	Edges []EdgeResponse `json:"edges"`
}

func NewWorkflowUIResponse(nodes []NodeResponse, edges []EdgeResponse) *WorkflowUIResponse {
	return &WorkflowUIResponse{
		Nodes: nodes,
		Edges: edges,
	}
}
