package pages

import (
	"testing"

	"github.com/Fumiya-Tahara/uecs-navi.git/internal/testdata"
	mock_repository "github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetM304IDs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mr := mock_repository.NewMockM304RepositoryInterface(ctrl)
	cdr := mock_repository.NewMockClimateDataRepositoryInterface(ctrl)
	dr := mock_repository.NewMockDeviceRepositoryInterface(ctrl)

	pus := NewPageUtilitiesService(mr, cdr, dr)

	m304s := testdata.GetTestM304s()
	wantM304IDs := []int{1, 2, 3}

	t.Run("return M304IDs", func(t *testing.T) {
		mr.EXPECT().GetAllM304s().Return(&m304s, nil)

		got, err := pus.GetM304IDs()
		assert.NoError(t, err, "Expected no error, got %v", err)

		assert.Equal(t, wantM304IDs, got, "Expected data %d, got %d", wantM304IDs, got)
	})

}

func TestGetClimateData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mr := mock_repository.NewMockM304RepositoryInterface(ctrl)
	cdr := mock_repository.NewMockClimateDataRepositoryInterface(ctrl)
	dr := mock_repository.NewMockDeviceRepositoryInterface(ctrl)

	pus := NewPageUtilitiesService(mr, cdr, dr)

	climateData := testdata.GetTestClimateData()

	t.Run("return climate data", func(t *testing.T) {
		cdr.EXPECT().GetAllClimateData().Return(&climateData, nil)

		got, err := pus.GetClimateData()
		assert.NoError(t, err, "Expected no error, got %v", err)

		assert.Equal(t, climateData, *got, "Expected data %d, got %d", climateData, *got)
	})

}

func TestGetDevices(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mr := mock_repository.NewMockM304RepositoryInterface(ctrl)
	cdr := mock_repository.NewMockClimateDataRepositoryInterface(ctrl)
	dr := mock_repository.NewMockDeviceRepositoryInterface(ctrl)

	pus := NewPageUtilitiesService(mr, cdr, dr)

	deviceMap := testdata.GetDevicesMap()

	t.Run("return devices connected to m304", func(t *testing.T) {
		for k, v := range deviceMap {
			dr.EXPECT().GetDevicesFromM304(k).Return(&v, nil)

			got, err := pus.GetDevices(k)
			assert.NoError(t, err, "Expected no error, got %v", err)
			assert.Equal(t, deviceMap[k], *got, "Expected data %d, got %d", deviceMap[k], *got)

		}
	})
}
