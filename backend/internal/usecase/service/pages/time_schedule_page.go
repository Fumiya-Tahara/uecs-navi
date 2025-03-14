package pages

import (
	"log"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/interfaces"
)

type TimeSchedulePageService struct {
	conditionRepository       interfaces.ConditionRepositoryInterface
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

// users table作成後はuser_idから各M304に登録されているtime_scheduleを全取得する
func (tsps TimeSchedulePageService) GetTimeSchedule(m304ID int) (*domain.TimeSchedule, error) {
	timeSchedule, err := tsps.timeScheduleRepository.GetTimeScheduleFromM304(m304ID)
	if err != nil {
		log.Printf("Error getting time schedule: %v", err)
		return nil, err
	}

	timeScheduleRows, err := tsps.timeScheduleRowRepository.GetTimeScheduleRowsFromTimeSchedule(timeSchedule.ID)
	if err != nil {
		log.Printf("Error getting time schedule rows: %v", err)
		return nil, err
	}

	timeSchedule.Rows = *timeScheduleRows

	return timeSchedule, nil
}

func (tsps TimeSchedulePageService) CreateAndBuildTimeSchedule(timeSchedule domain.TimeSchedule) error {
	_, err := tsps.timeScheduleRepository.CreateTimeSchedule(timeSchedule.M304ID)
	if err != nil {
		log.Printf("Error creating time schedule: %v", err)
		return err
	}

	for _, v := range timeSchedule.Rows {
		_, err := tsps.timeScheduleRowRepository.CreateTimeScheduleRow(v)
		if err != nil {
			log.Printf("Error creating time schedule row: %v", err)

			// エラーハンドリング要修正
			return err
		}

		_, err = tsps.conditionRepository.CreateCondition(&v.Condition)
		if err != nil {
			log.Printf("Error creating condition: %v", err)

			return err
		}
	}

	// buildするときの処理を追加

	return nil
}

func (tsps TimeSchedulePageService) GetWorkflows(m304ID int) (*[]domain.Workflow, error) {
	workflows, err := tsps.workflowRepository.GetWorkflowsFromM304(m304ID)
	if err != nil {
		log.Printf("Error getting workflows: %v", err)
		return nil, err
	}

	return workflows, nil
}

func (tsps TimeSchedulePageService) UpdateAndBuildTimeSchedule(timeSchedule domain.TimeSchedule) error {
	if err := tsps.timeScheduleRepository.UpdateTimeSchedule(timeSchedule); err != nil {
		log.Printf("Error updating time schedule: %v", err)
		return err
	}

	timeScheduleRows := timeSchedule.Rows
	for _, timeScheduleRow := range timeScheduleRows {
		if err := tsps.timeScheduleRowRepository.UpdateTimeScheduleRow(timeScheduleRow); err != nil {
			log.Printf("Error updating time schedule row: %v", err)

			return err
		}

		if err := tsps.conditionRepository.UpdateCondition(timeScheduleRow.Condition); err != nil {
			log.Printf("Error updating condition: %v", err)

			return err
		}
	}

	// buildするときの処理を追加

	return nil
}
