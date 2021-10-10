package data

import "errors"

type PlaceInfo struct {
	ID string
	Address string
	Phone string
	HoursOfWork string
	CoordsLat string
	CoordsLon string
	IsOpen bool
}

func(pi *PlaceInfo) Validate() error {
	if len(pi.Address) == 0 {
		return errors.New("empty address")
	}
	return nil
}

type PlaceInfoRepository interface {
	Store(info PlaceInfo) error
	GetByID(id string) (PlaceInfo, error)
	GetByAddressLike(address string) ([]PlaceInfo, error)
	IsOpen(id string) (bool, error)
}
