package dto

type WorkflowResponse struct {
	ID     int    `json:"id"`
	M304ID int    `json:"m304_id"`
	Name   string `json:"name"`
}

func NewWorkflowResponse(id, m304ID int, name string) *WorkflowResponse {
	return &WorkflowResponse{
		ID:     id,
		M304ID: m304ID,
		Name:   name,
	}
}
