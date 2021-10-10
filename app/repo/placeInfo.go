package repo

import (
	"CoffeLand/app/core/data"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type PlaceInfo struct {
	db gorm.DB
}

func (p PlaceInfo) Store(info data.PlaceInfo) error {
	_, err := p.GetByID(info.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return p.db.Create(&info).Error
		}
		return errors.WithStack(err)
	}

	return p.db.Save(info).Error
}

func (p PlaceInfo) GetByID(id string) (data.PlaceInfo, error) {
	var info data.PlaceInfo
	result := p.db.Where("id = ?", id).Find(&info)
	if result.Error != nil {
		return data.PlaceInfo{}, result.Error
	}
	if result.RowsAffected == 0 {
		return data.PlaceInfo{}, gorm.ErrRecordNotFound
	}

	return info, nil
}

func (p PlaceInfo) GetByAddressLike(address string) ([]data.PlaceInfo, error) {
	var info []data.PlaceInfo
	result := p.db.Where("address LIKE ?", address + "%").Find(&info)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return []data.PlaceInfo{}, gorm.ErrRecordNotFound
	}
	return info, nil
}

func (p PlaceInfo) IsOpen(id string) (bool, error) {
	info, err := p.GetByID(id)
	if err != nil {
		return false, errors.WithStack(err)
	}

	return info.IsOpen, nil
}



