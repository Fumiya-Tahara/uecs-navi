package interfaces

import "github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"

type (
	// m304 interfaces
	TimeScheduleM304Interface interface {
		SendTimeSchedule(timeScheduleRow domain.TimeScheduleRow) error
	}

	WorkflowM304Interface interface {
		SendWorkflow(workflowOperation domain.WorkflowOperation) error
	}
)
