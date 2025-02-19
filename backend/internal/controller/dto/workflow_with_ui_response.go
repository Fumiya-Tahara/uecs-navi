package dto

type WorkflowWithUIResponse struct {
	Workflow   Workflow           `json:"workflow"`
	WorkflowUI WorkflowUIResponse `json:"workflowUI"`
}

func NewWorkflowWithUIResponse(workflow Workflow, workflowUI WorkflowUIResponse) *WorkflowWithUIResponse {
	return &WorkflowWithUIResponse{
		Workflow:   workflow,
		WorkflowUI: workflowUI,
	}
}
