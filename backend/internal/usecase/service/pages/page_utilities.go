package pages

import (
	"log"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/interfaces"
)

type PageUtilitiesService struct {
	m304Repository        interfaces.M304RepositoryInterface
	climateDataRepository interfaces.ClimateDataRepositoryInterface
}

func NewPageUtilitiesService(mr interfaces.M304RepositoryInterface, cdr interfaces.ClimateDataRepositoryInterface) *PageUtilitiesService {
	return &PageUtilitiesService{
		m304Repository:        mr,
		climateDataRepository: cdr,
	}
}

func (pus PageUtilitiesService) GetM304IDs() ([]int, error) {
	m304s, err := pus.m304Repository.GetAllM304s()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	m304IDs := make([]int, len(*m304s))
	for i, m304 := range *m304s {
		m304IDs[i] = m304.ID
	}

	return m304IDs, nil
}

func (pus PageUtilitiesService) GetClimateData() (*[]domain.ClimateData, error) {
	climateData, err := pus.climateDataRepository.GetAllClimateData()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return climateData, nil
}
