package interfaces

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
)

type (
	// Repository interfaces
	ClimateDataRepositoryInterface interface {
		GetAllClimateData() (*[]domain.ClimateData, error)
		GetClimateDataFromID(ID int) (*domain.ClimateData, error)
	}

	ComparisonOperatorRepositoryInterface interface {
		GetAllComparisonOperators() (*[]domain.ComparisonOperator, error)
		GetAllComparisonOperatorFromID(ID int) *domain.ComparisonOperator
	}

	ConditionRepositoryInterface interface {
		CreateCondition(newCondition *domain.Condition) (int, error)
		GetConditionFromTimeScheduleRow(timeScheduleRowID int) (*domain.Condition, error)
		UpdateCondition(condition domain.Condition) error
	}

	DeviceRepositoryInterface interface {
		CreateDevice(newDevice domain.Device) (int, error)
		GetDevicesFromM304(m304ID int) (*[]domain.Device, error)
	}

	EdgeRepositoryInterface interface {
		CreateEdge(newEdge domain.Edge) (int, error)
		GetEdgesFromWorkflow(workflowID int) (*[]domain.Edge, error)
		UpdateEdge(edge domain.Edge) error
	}

	HouseRepositoryInterface interface {
		CreateHouse(name string) (int, error)
		GetAllHouses() (*[]domain.House, error)
	}

	M304RepositoryInterface interface {
		CreateM304(newM304 domain.M304) (int, error)
		GetAllM304s() (*[]domain.M304, error)
		GetM304FromID(ID int) (*domain.M304, error)
	}

	NodeRepositoryInterface interface {
		CreateNode(newNode domain.Node) (int, error)
		GetNodesFromWorkflow(workflowID int) (*[]domain.Node, error)
		UpdateNode(node domain.Node) error
	}

	TimeScheduleRowRepositoryInterface interface {
		CreateTimeScheduleRow(newTimeScheduleRow domain.TimeScheduleRow) (int, error)
		GetTimeScheduleRowsFromTimeSchedule(timeScheduleID int) (*[]domain.TimeScheduleRow, error)
		UpdateTimeScheduleRow(timeScheduleRow domain.TimeScheduleRow) error
	}

	TimeScheduleRepositoryInterface interface {
		CreateTimeSchedule(m304ID int) (int, error)
		GetTimeScheduleFromM304(m304ID int) (*domain.TimeSchedule, error)
		UpdateTimeSchedule(timeSchedule domain.TimeSchedule) error
	}

	WorkflowOperationRepositoryInterface interface {
		CreateWorkflowOperation(newWorkflowOperation domain.WorkflowOperation) (int, error)
		GetWorkflowOperationsFromWorkflow(workflowID int) (*domain.WorkflowOperation, error)
		UpdateWorkflowOperation(workflowOperation domain.WorkflowOperation) error
	}

	WorkflowRepositoryInterface interface {
		CreateWorkflow(newWorkflow domain.Workflow) (int, error)
		GetWorkflowsFromM304(m304ID int) (*[]domain.Workflow, error)
		UpdateWorkflow(workflow domain.Workflow) error
		DeleteWorkflow(workflowID int) error
	}
)
