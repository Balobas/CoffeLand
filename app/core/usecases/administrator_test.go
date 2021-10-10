package usecases

import (
	"CoffeLand/app/core/data"
	"fmt"
	"testing"
)

type AdminRepoMock struct {

}

func (a AdminRepoMock) Store(data.Administrator) error {
	return nil
}

func (a AdminRepoMock) GetByID(string) (data.Administrator, error) {
	return data.Administrator{
		ID:         "123",
		FirstName:  "",
		LastName:   "",
		Patronymic: "",
		Password:   "",
		Access:     data.AdministratorAccessMiddle,
		Hash:       nil,
	}, nil
}

func (a AdminRepoMock) GetByFIO(firstName string, lastName string, patronymic string) ([]data.Administrator, error) {
	return []data.Administrator{}, nil
}

var (
	CoreMock = data.Core{
		AdminRepo:    AdminRepoMock{},
		DiscountRepo: nil,
		ProductRepo:  nil,
		BlogPostRepo: nil,
	}
)


func TestNewAdministratorUsecases(t *testing.T) {
	adminUsecases, err := NewAdministratorUsecases(&CoreMock, "123")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	fmt.Println(adminUsecases.Admin)
}

func TestAdministratorUsecases_PutAdministrator(t *testing.T) {
	adminUsecases, err := NewAdministratorUsecases(&CoreMock, "123")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	_, err = adminUsecases.PutAdministrator(data.Administrator{
		ID:         "",
		FirstName:  "vasya",
		LastName:   "petrov",
		Patronymic: "petrovich",
		Password:   "",
		Access:     data.AdministratorAccessLow,
		Hash:       nil,
	})

	if err != nil {
		t.Log(err)
	} else {
		t.Errorf("expected error")
		t.FailNow()
	}
}


