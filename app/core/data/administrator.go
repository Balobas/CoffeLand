package data

import (
	"github.com/pkg/errors"
)

type Administrator struct {
	ID string `gorm:"primarykey"`
	FirstName string
	LastName string
	Patronymic string
	Password string
	Access AdministratorAccess
	Hash []byte
}


var (
	AccessError = errors.New("AccessError")
)

func(a *Administrator) HasProductAccess() error {
	if a.Access != AdministratorAccessMiddle && a.Access != AdministratorAccessHigh {
		return AccessError
	}
	return nil
}

func(a *Administrator) HasDiscountAccess() error {
	if a.Access != AdministratorAccessMiddle && a.Access != AdministratorAccessHigh {
		return AccessError
	}
	return nil
}

func(a *Administrator) HasAddAdministratorAccess() error {
	if a.Access != AdministratorAccessHigh {
		return AccessError
	}
	return nil
}

func(a *Administrator) HasAddBlogPostsAccess() error {
	if a.Access != AdministratorAccessMiddle && a.Access != AdministratorAccessHigh {
		return AccessError
	}
	return nil
}

func(a *Administrator) Validate() error {
	//hashParams := []byte(a.FirstName)
	//hashParams = append(hashParams, []byte(a.LastName)...)
	//hashParams = append(hashParams, []byte(a.Patronymic)...)
	//hashParams = append(hashParams, []byte(a.Password)...)
	//md5.New().Sum(hashParams)
	//if string(hashParams) != string(a.Hash) {
	//	return errors.New("hash error")
	//}
	return nil
}

type AdministratorRepository interface {
	Store(Administrator) error
	GetByID(string) (Administrator, error)
	GetByFIO(firstName string, lastName string, patronymic string) ([]Administrator, error)
}
