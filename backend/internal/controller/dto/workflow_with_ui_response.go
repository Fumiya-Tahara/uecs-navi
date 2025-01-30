package dto

type WorkflowWithUIResponse struct {
	Workflow   WorkflowResponse   `json:"workflow"`
	WorkflowUI WorkflowUIResponse `json:"workflowUI"`
}

func NewWorkflowWithUIResponse(workflow WorkflowResponse, workflowUI WorkflowUIResponse) *WorkflowWithUIResponse {
	return &WorkflowWithUIResponse{
		Workflow:   workflow,
		WorkflowUI: workflowUI,
	}
}
