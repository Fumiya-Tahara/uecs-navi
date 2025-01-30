package dto

type WorkflowWithUIRequest struct {
	Workflow   Workflow          `json:"workflow"`
	WorkflowUI WorkflowUIRequest `json:"workflow_ui"`
	Relays     Relays            `json:"relays"`
}

func NewWorkflowWithUIRequest(workflow Workflow, workflowUI WorkflowUIRequest, relays Relays) *WorkflowWithUIRequest {
	return &WorkflowWithUIRequest{
		Workflow:   workflow,
		WorkflowUI: workflowUI,
		Relays:     relays,
	}
}
