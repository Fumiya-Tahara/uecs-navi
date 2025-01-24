package repository

import (
	"context"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/utils"
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
		Relay1:     utils.PointerToNullBool(newWorkflowOperation.Relay1),
		Relay2:     utils.PointerToNullBool(newWorkflowOperation.Relay2),
		Relay3:     utils.PointerToNullBool(newWorkflowOperation.Relay3),
		Relay4:     utils.PointerToNullBool(newWorkflowOperation.Relay4),
		Relay5:     utils.PointerToNullBool(newWorkflowOperation.Relay5),
		Relay6:     utils.PointerToNullBool(newWorkflowOperation.Relay6),
		Relay7:     utils.PointerToNullBool(newWorkflowOperation.Relay7),
		Relay8:     utils.PointerToNullBool(newWorkflowOperation.Relay8),
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
		utils.NullBoolToPointer(workflowOperationRow.Relay1),
		utils.NullBoolToPointer(workflowOperationRow.Relay2),
		utils.NullBoolToPointer(workflowOperationRow.Relay3),
		utils.NullBoolToPointer(workflowOperationRow.Relay4),
		utils.NullBoolToPointer(workflowOperationRow.Relay5),
		utils.NullBoolToPointer(workflowOperationRow.Relay6),
		utils.NullBoolToPointer(workflowOperationRow.Relay7),
		utils.NullBoolToPointer(workflowOperationRow.Relay8),
	)

	return workflowOperations, nil
}
