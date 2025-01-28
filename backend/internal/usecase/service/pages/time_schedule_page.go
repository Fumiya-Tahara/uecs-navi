package pages

import (
	"log"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/interfaces"
)

type TimeSchedulePageService struct {
	timeScheduleRepository    interfaces.TimeScheduleRepositoryInterface
	timeScheduleRowRepository interfaces.TimeScheduleRowRepositoryInterface
	workflowRepository        interfaces.WorkflowRepositoryInterface
}

func NewTimeSchedulePageService(tsr interfaces.TimeScheduleRepositoryInterface, tsrr interfaces.TimeScheduleRowRepositoryInterface, wr interfaces.WorkflowRepositoryInterface) *TimeSchedulePageService {
	return &TimeSchedulePageService{
		timeScheduleRepository:    tsr,
		timeScheduleRowRepository: tsrr,
		workflowRepository:        wr,
	}
}

func (tsps TimeSchedulePageService) GetTimeSchedule(m304ID int) (*domain.TimeSchedule, error) {
	timeSchedule, err := tsps.timeScheduleRepository.GetTimeScheduleFromM304(m304ID)
	if err != nil {
		log.Printf("Error getting time schedule: %v", err)
	}

	timeScheduleRows, err := tsps.timeScheduleRowRepository.GetTimeScheduleRowsFromTimeSchedule(timeSchedule.ID)
	if err != nil {
		log.Printf("Error getting time schedule rows: %v", err)
	}

	timeSchedule.Rows = *timeScheduleRows

	return timeSchedule, nil
}

func (tsps TimeSchedulePageService) CreateAndBuildTimeSchedule(timeSchedule domain.TimeSchedule) error {
	_, err := tsps.timeScheduleRepository.CreateTimeSchedule(timeSchedule.M304ID)
	if err != nil {
		log.Printf("Error creating time schedule: %v", err)
	}

	for _, v := range timeSchedule.Rows {
		_, err := tsps.timeScheduleRowRepository.CreateTimeScheduleRow(v)
		if err != nil {
			log.Printf("Error creating time schedule row: %v", err)
		}
	}

	return nil
}

func (tsps TimeSchedulePageService) GetWorkflows(m304ID int) (*[]domain.Workflow, error) {
	workflows, err := tsps.workflowRepository.GetWorkflowsFromM304(m304ID)
	if err != nil {
		log.Printf("Error getting workflows: %v", err)
	}

	return workflows, nil
}
