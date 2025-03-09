package controller

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/db/mysql"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/infrastructure/orm/mysqlc"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/repository"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/service/pages"
)

func InitHandler() (*Handler, error) {
	db, err := mysql.ConnectDB()
	if err != nil {
		return nil, err
	}

	query := mysqlc.New(db)

	wr := repository.NewWorkflowRepository(query)
	wor := repository.NewWorkflowOperationRepository(query)
	nr := repository.NewNodeRepository(query)
	er := repository.NewEdgeRepository(query)
	tsr := repository.NewTimeScheduleRepository(query)
	tsrr := repository.NewTimeScheduleRowRepository(query)
	mr := repository.NewM304Repository(query)
	cdr := repository.NewClimateDataRepository(query)

	wps := pages.NewWorkflowPageService(wr, wor, nr, er)
	tsps := pages.NewTimeSchedulePageService(tsr, tsrr, wr)
	pus := pages.NewPageUtilitiesService(mr, cdr)

	h := NewHandler(*wps, *tsps, *pus)

	return h, nil
}
