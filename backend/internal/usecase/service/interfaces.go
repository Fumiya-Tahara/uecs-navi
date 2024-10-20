package service

import (
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/domain"
	"github.com/Fumiya-Tahara/uecs-navi.git/internal/usecase/repository"
)

type (
	// Repository interfaces
	DeviceRepositoryInterface interface {
		CreateDevice(newDevice domain.Device) (int64, error)
		GetDevicesFromHouse(houseID int) ([]*domain.Device, error)
		GetJoinedDevicesFromHouse(houseID int) ([]*repository.JoinedDevice, error)
	}

	HouseRepositoryInterface interface {
		CreateHouse(name string) (int64, error)
		GetAllHouses() ([]*domain.House, error)
	}

	ClimateDataRepositoryInterface interface {
		GetAllClimateData() ([]*domain.ClimateData, error)
	}
)
