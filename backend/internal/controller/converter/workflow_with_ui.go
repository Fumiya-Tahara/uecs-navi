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

func ToDomainWorkflow(workflowReq dto.WorkflowWithUIRequest) *domain.Workflow {
	relays := workflowReq.Relays
	undefinedWorkflowID := 0
	workflowOperation := domain.NewWorkflowOperation(
		undefinedWorkflowID,
		&relays.Relay1,
		&relays.Relay2,
		&relays.Relay3,
		&relays.Relay4,
		&relays.Relay5,
		&relays.Relay6,
		&relays.Relay7,
		&relays.Relay8,
	)

	nodes := ToDomainNode(workflowReq.WorkflowUI.Nodes)
	edges := ToDomainEdge(workflowReq.WorkflowUI.Edges)

	workflow := domain.NewWorkflow(workflowReq.Workflow.M304ID, workflowReq.Workflow.Name, *workflowOperation, *nodes, *edges)

	return workflow
}
