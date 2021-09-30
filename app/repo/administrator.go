package repo

import (
	"CoffeLand/app/core/data"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
)

type Administrator struct {
	db gorm.DB
}

func NewAdministratorRepo(db gorm.DB) data.AdministratorRepository {
	return Administrator{db:db}
}

func (a Administrator) Store(administrator data.Administrator) error {
	_, err := a.GetByID(administrator.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return a.db.Create(&administrator).Error
		}
		return errors.WithStack(err)
	}

	return a.db.Save(administrator).Error
}

func (a Administrator) GetByID(id string) (data.Administrator, error) {
	var administrator data.Administrator
	result := a.db.Where("id = ?", id).Find(&administrator)
	if result.Error != nil {
		return data.Administrator{}, result.Error
	}
	if result.RowsAffected == 0 {
		return data.Administrator{}, gorm.ErrRecordNotFound
	}

	return administrator, nil
}

func (a Administrator) GetByFIO(firstName string, lastName string, patronymic string) ([]data.Administrator, error) {
	var whereCases []string
	var whereParams []string
	if len(patronymic) != 0 {
		whereCases = append(whereCases, "patronymic LIKE ?")
		whereParams = append(whereParams, patronymic + "%")
	}
	if len(lastName) != 0 {
		whereCases = append(whereCases, "last_name LIKE ?")
		whereParams = append(whereParams, lastName + "%")
	}
	if len(firstName) != 0 {
		whereCases = append(whereCases, "first_name LIKE ?")
		whereParams = append(whereParams, firstName + "%")
	}

	var (
		admins []data.Administrator
		result *gorm.DB
	)

	switch len(whereCases) {
	case 0:
		result := a.db.Find(&admins)
		if result.Error != nil {
			return []data.Administrator{}, result.Error
		}
		if result.RowsAffected == 0 {
			return []data.Administrator{}, gorm.ErrRecordNotFound
		}
		return admins, nil
	case 1:
		result = a.db.Where(strings.Join(whereCases, " AND "), whereParams[0]).Find(&admins)
	case 2:
		result = a.db.Where(strings.Join(whereCases, " AND "), whereParams[0], whereParams[1]).Find(&admins)
	default:
		result = a.db.Where(strings.Join(whereCases, " AND "), whereParams[0], whereParams[1], whereParams[2]).Find(&admins)
	}

	if result.Error != nil {
		return []data.Administrator{}, result.Error
	}
	if result.RowsAffected == 0 {
		return []data.Administrator{}, gorm.ErrRecordNotFound
	}
	return admins, nil
}


