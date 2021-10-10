package usecases

import "CoffeLand/app/core/data"

type PlaceInfoRepoMock struct {
	StoreError error
	GetByIDInfo data.PlaceInfo
	GetByIDError error
	GetByAddressInfo []data.PlaceInfo
	GetByAddressError error
	IsOpenFlag bool
	IsOpenError error
}

func (p PlaceInfoRepoMock) Store(_ data.PlaceInfo) error {
	return p.StoreError
}

func (p PlaceInfoRepoMock) GetByID(_ string) (data.PlaceInfo, error) {
	return p.GetByIDInfo, p.GetByIDError
}

func (p PlaceInfoRepoMock) GetByAddressLike(_ string) ([]data.PlaceInfo, error) {
	return p.GetByAddressInfo, p.GetByAddressError
}

func (p PlaceInfoRepoMock) IsOpen(_ string) (bool, error) {
	return p.IsOpenFlag, p.IsOpenError
}

