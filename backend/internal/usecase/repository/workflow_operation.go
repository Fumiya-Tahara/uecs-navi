package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
)

type WorkflowOperationRepository struct {
	queries *mysqlc.Queries
}

func NewWorkflowOperationRepository(queries *mysqlc.Queries) *WorkflowOperationRepository {
	return &WorkflowOperationRepository{
		queries: queries,
	}
}

func (wor *WorkflowOperationRepository) CreateWorkflowOperation(newWorkflowOperation domain.WorkflowOperation) (int, error) {
	ctx := context.Background()

	arg := mysqlc.CreateWorkflowOperationParams{
		WorkflowID: int32(newWorkflowOperation.WorkflowID),
		Relay1:     PointerToNullBool(newWorkflowOperation.Relay1),
		Relay2:     PointerToNullBool(newWorkflowOperation.Relay2),
		Relay3:     PointerToNullBool(newWorkflowOperation.Relay3),
		Relay4:     PointerToNullBool(newWorkflowOperation.Relay4),
		Relay5:     PointerToNullBool(newWorkflowOperation.Relay5),
		Relay6:     PointerToNullBool(newWorkflowOperation.Relay6),
		Relay7:     PointerToNullBool(newWorkflowOperation.Relay7),
		Relay8:     PointerToNullBool(newWorkflowOperation.Relay8),
	}

	id, err := wor.queries.CreateWorkflowOperation(ctx, arg)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (wor *WorkflowOperationRepository) GetWorkflowOperationsFromWorkflow(workflowID int) (*domain.WorkflowOperation, error) {
	ctx := context.Background()

	workflowOperationRow, err := wor.queries.GetWorkflowOperationsFromWorkflow(ctx, int32(workflowID))
	if err != nil {
		return nil, err
	}

	workflowOperations := domain.NewWorkflowOperationWithID(
		int(workflowOperationRow.ID),
		int(workflowOperationRow.WorkflowID),
		NullBoolToPointer(workflowOperationRow.Relay1),
		NullBoolToPointer(workflowOperationRow.Relay2),
		NullBoolToPointer(workflowOperationRow.Relay3),
		NullBoolToPointer(workflowOperationRow.Relay4),
		NullBoolToPointer(workflowOperationRow.Relay5),
		NullBoolToPointer(workflowOperationRow.Relay6),
		NullBoolToPointer(workflowOperationRow.Relay7),
		NullBoolToPointer(workflowOperationRow.Relay8),
	)

	return workflowOperations, nil
}
