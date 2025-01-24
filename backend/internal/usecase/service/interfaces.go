package service

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
)

type (
	// Repository interfaces
	ClimateDataRepositoryInterface interface {
		GetAllClimateData() ([]*domain.ClimateData, error)
		GetClimateDataFromID(ID int) (*domain.ClimateData, error)
	}

	ComparisonOperatorRepositoryInterface interface {
		GetAllComparisonOperators() ([]*domain.ComparisonOperator, error)
		GetAllComparisonOperatorFromID(ID int) *domain.ComparisonOperator
	}

	ConditionRepositoryInterface interface {
		CreateCondition(newCondition *domain.Condition) (int, error)
		GetConditionFromTimeScheduleRow(timeScheduleRowID int) (*domain.Condition, error)
	}

	DeviceRepositoryInterface interface {
		CreateDevice(newDevice domain.Device) (int, error)
		GetDevicesFromM304(m304ID int) ([]*domain.Device, error)
	}

	EdgeRepositoryInterface interface {
		CreateEdge(newEdge domain.Edge) (int, error)
		GetEdgesFromWorkflow(workflowID int) ([]*domain.Edge, error)
	}

	HouseRepositoryInterface interface {
		CreateHouse(name string) (int, error)
		GetAllHouses() ([]*domain.House, error)
	}

	M304RepositoryInterface interface {
		CreateM304(newM304 domain.M304) (int, error)
		GetM304FromID(ID int) (*domain.M304, error)
	}

	NodeRepositoryInterface interface {
		CreateNode(newNode domain.Node) (int, error)
		GetNodesFromWorkflow(workflowID int) ([]*domain.Node, error)
	}

	TimeScheduleRowRepositoryInterface interface {
		CreateTimeScheduleRow(newTimeScheduleRow domain.TimeScheduleRow) (int, error)
		GetTimeScheduleRowsFromTimeSchedule(timeScheduleID int) ([]*domain.TimeScheduleRow, error)
	}

	TimeScheduleRepositoryInterface interface {
		CreateTimeSchedule(m304ID int) (int, error)
		GetTimeScheduleFromM304(m304ID int, rows []domain.TimeScheduleRow) (*domain.TimeSchedule, error)
	}

	WorkflowOperationRepositoryInterface interface {
		CreateWorkflowOperation(newWorkflowOperation domain.WorkflowOperation) (int, error)
		GetWorkflowOperationsFromWorkflow(workflowID int) (*domain.WorkflowOperation, error)
	}

	WorkflowRepositoryInterface interface {
		CreateWorkflow(newWorkflow domain.Workflow) (int, error)
		GetWorkflowsFromM304(m304ID int) ([]*domain.Workflow, error)
	}
)
