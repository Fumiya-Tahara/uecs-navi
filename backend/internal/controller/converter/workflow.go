package converter

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/controller/dto"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
)

func ToDtoWorkflow(workflow domain.Workflow) *dto.Workflow {
	dtoWorkflow := dto.NewWorkflow(
		workflow.ID,
		workflow.M304ID,
		workflow.Name,
	)

	return dtoWorkflow
}
