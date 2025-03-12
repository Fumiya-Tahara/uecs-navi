package testdata

import "github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"

func GetTestClimateData() []domain.ClimateData {
	climateData := []domain.ClimateData{
		*domain.NewClimateData(1, "気温", "℃"),
		*domain.NewClimateData(2, "湿度", "%"),
		*domain.NewClimateData(3, "二酸化炭素", "ppm"),
		*domain.NewClimateData(4, "照度", "lx"),
	}

	return climateData
}
