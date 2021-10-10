package usecases

import (
	"CoffeLand/app/core/data"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type PlaceInfoUsecases struct {
	*data.Core
}

func NewPlaceInfoUsecases(core *data.Core) *PlaceInfoUsecases {
	return &PlaceInfoUsecases{core}
}

func(pic *PlaceInfoUsecases) put(info data.PlaceInfo) (string, error) {
	if err := info.Validate(); err != nil {
		return "", errors.WithStack(err)
	}

	if len(info.ID) == 0 {
		info.ID = uuid.New().String()
	}

	if err := pic.PlaceInfoRepo.Store(info); err != nil {
		return info.ID, err
	}

	return info.ID, nil
}

func(pic *PlaceInfoUsecases) GetByID(id string) (data.PlaceInfo, error) {
	return pic.PlaceInfoRepo.GetByID(id)
}

func(pic *PlaceInfoUsecases) IsOpen(id string) (bool, error) {
	return pic.PlaceInfoRepo.IsOpen(id)
}

func(pic *PlaceInfoUsecases) GetByAddressLike(address string) ([]data.PlaceInfo, error) {
	return pic.PlaceInfoRepo.GetByAddressLike(address)
}
