package converter

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/dto"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
)

func ToDtoWorkflowWithUI(workflows []domain.Workflow) *[]dto.WorkflowWithUIResponse {
	workflowsRes := make([]dto.WorkflowWithUIResponse, len(workflows))
	for i, workflow := range workflows {
		workflowRes := dto.NewWorkflow(workflow.ID, workflow.M304ID, workflow.Name)

		nodeRes := ToDtoNodeResponse(workflow.Node)
		edgeRes := ToDtoEdgeResponse(workflow.Edge)
		workflowUI := dto.NewWorkflowUIResponse(*nodeRes, *edgeRes)

		workflowsRes[i] = *dto.NewWorkflowWithUIResponse(*workflowRes, *workflowUI)
	}

	return &workflowsRes
}
