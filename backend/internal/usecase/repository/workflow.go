package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type WorkflowRepository struct {
	queries *mysqlc.Queries
}

func NewWorkflowRepository(queries *mysqlc.Queries) *WorkflowRepository {
	return &WorkflowRepository{
		queries: queries,
	}
}

func (wr WorkflowRepository) CreateWorkflow(newWorkflow domain.Workflow) (int, error) {
	ctx := context.Background()

	arg := mysqlc.CreateWorkflowParams{
		M304ID: int32(newWorkflow.M304ID),
		Name:   newWorkflow.Name,
	}

	id, err := wr.queries.CreateWorkflow(ctx, arg)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (wr WorkflowRepository) GetWorkflowsFromM304(m304ID int) (*[]domain.Workflow, error) {
	ctx := context.Background()

	workflowRows, err := wr.queries.GetWorkflowsFromM304(ctx, int32(m304ID))
	if err != nil {
		return nil, err
	}

	workflows := make([]domain.Workflow, len(workflowRows))
	for i, v := range workflowRows {
		workflows[i] = *domain.NewWorkflowWithBasicInfo(
			int(v.ID),
			int(v.M304ID),
			v.Name,
		)
	}

	return &workflows, nil
}

func (wr WorkflowRepository) UpdateWorkflow(workflow domain.Workflow) error {
	ctx := context.Background()

	arg := mysqlc.UpdateWorkflowParams{
		M304ID: int32(workflow.M304ID),
		Name:   workflow.Name,
		ID:     int32(workflow.ID),
	}

	if err := wr.queries.UpdateWorkflow(ctx, arg); err != nil {
		return err
	}

	return nil
}

func (wr WorkflowRepository) DeleteWorkflow(workflowID int) error {
	ctx := context.Background()

	if err := wr.queries.DeleteWorkflow(ctx, int32(workflowID)); err != nil {
		return err
	}

	return nil
}
